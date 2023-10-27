package main

import (
	"fmt"
	"io/fs"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/go-vgo/robotgo"
	"github.com/joho/godotenv"
)

type PlayerPos struct {
	x  string
	y  string
	z  string
	rx string
	rz string
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	launchGame()
	createNewWorld()
	setupScreenshot()
	for i := 0; i < 5; i++ {
		takeRandomScreenshot()
		postScreenshotToSocialMedia()
	}
	quitGame()
	cleanup()
}

func launchGame() {
	fmt.Println("Starting Minecraft...")
	cmd := exec.Command(
		"/Users/alyx/Minecraft/Install/runtime/java-runtime-gamma/mac-os-arm64/java-runtime-gamma/jre.bundle/Contents/Home/bin/java",
		"-XstartOnFirstThread",
		"-Djava.library.path=/Users/alyx/Minecraft/Install/bin/5fd53412750b559c393aee4a091a4682d39fbecb",
		"-Djna.tmpdir=/Users/alyx/Minecraft/Install/bin/5fd53412750b559c393aee4a091a4682d39fbecb",
		"-Dorg.lwjgl.system.SharedLibraryExtractPath=/Users/alyx/Minecraft/Install/bin/5fd53412750b559c393aee4a091a4682d39fbecb",
		"-Dio.netty.native.workdir=/Users/alyx/Minecraft/Install/bin/5fd53412750b559c393aee4a091a4682d39fbecb",
		"-Dminecraft.launcher.brand=minecraft-launcher",
		"-Dminecraft.launcher.version=2.12.17",
		"-cp", "/Users/alyx/Minecraft/Install/libraries/net/fabricmc/tiny-mappings-parser/0.3.0+build.17/tiny-mappings-parser-0.3.0+build.17.jar:/Users/alyx/Minecraft/Install/libraries/net/fabricmc/sponge-mixin/0.12.5+mixin.0.8.5/sponge-mixin-0.12.5+mixin.0.8.5.jar:/Users/alyx/Minecraft/Install/libraries/net/fabricmc/tiny-remapper/0.8.2/tiny-remapper-0.8.2.jar:/Users/alyx/Minecraft/Install/libraries/net/fabricmc/access-widener/2.1.0/access-widener-2.1.0.jar:/Users/alyx/Minecraft/Install/libraries/org/ow2/asm/asm/9.5/asm-9.5.jar:/Users/alyx/Minecraft/Install/libraries/org/ow2/asm/asm-analysis/9.5/asm-analysis-9.5.jar:/Users/alyx/Minecraft/Install/libraries/org/ow2/asm/asm-commons/9.5/asm-commons-9.5.jar:/Users/alyx/Minecraft/Install/libraries/org/ow2/asm/asm-tree/9.5/asm-tree-9.5.jar:/Users/alyx/Minecraft/Install/libraries/org/ow2/asm/asm-util/9.5/asm-util-9.5.jar:/Users/alyx/Minecraft/Install/libraries/net/fabricmc/intermediary/1.20.2/intermediary-1.20.2.jar:/Users/alyx/Minecraft/Install/libraries/net/fabricmc/fabric-loader/0.14.22/fabric-loader-0.14.22.jar:/Users/alyx/Minecraft/Install/libraries/ca/weblite/java-objc-bridge/1.1/java-objc-bridge-1.1.jar:/Users/alyx/Minecraft/Install/libraries/com/github/oshi/oshi-core/6.4.5/oshi-core-6.4.5.jar:/Users/alyx/Minecraft/Install/libraries/com/google/code/gson/gson/2.10.1/gson-2.10.1.jar:/Users/alyx/Minecraft/Install/libraries/com/google/guava/failureaccess/1.0.1/failureaccess-1.0.1.jar:/Users/alyx/Minecraft/Install/libraries/com/google/guava/guava/32.1.2-jre/guava-32.1.2-jre.jar:/Users/alyx/Minecraft/Install/libraries/com/ibm/icu/icu4j/73.2/icu4j-73.2.jar:/Users/alyx/Minecraft/Install/libraries/com/mojang/authlib/5.0.47/authlib-5.0.47.jar:/Users/alyx/Minecraft/Install/libraries/com/mojang/blocklist/1.0.10/blocklist-1.0.10.jar:/Users/alyx/Minecraft/Install/libraries/com/mojang/brigadier/1.1.8/brigadier-1.1.8.jar:/Users/alyx/Minecraft/Install/libraries/com/mojang/datafixerupper/6.0.8/datafixerupper-6.0.8.jar:/Users/alyx/Minecraft/Install/libraries/com/mojang/logging/1.1.1/logging-1.1.1.jar:/Users/alyx/Minecraft/Install/libraries/com/mojang/patchy/2.2.10/patchy-2.2.10.jar:/Users/alyx/Minecraft/Install/libraries/com/mojang/text2speech/1.17.9/text2speech-1.17.9.jar:/Users/alyx/Minecraft/Install/libraries/commons-codec/commons-codec/1.16.0/commons-codec-1.16.0.jar:/Users/alyx/Minecraft/Install/libraries/commons-io/commons-io/2.13.0/commons-io-2.13.0.jar:/Users/alyx/Minecraft/Install/libraries/commons-logging/commons-logging/1.2/commons-logging-1.2.jar:/Users/alyx/Minecraft/Install/libraries/io/netty/netty-buffer/4.1.97.Final/netty-buffer-4.1.97.Final.jar:/Users/alyx/Minecraft/Install/libraries/io/netty/netty-codec/4.1.97.Final/netty-codec-4.1.97.Final.jar:/Users/alyx/Minecraft/Install/libraries/io/netty/netty-common/4.1.97.Final/netty-common-4.1.97.Final.jar:/Users/alyx/Minecraft/Install/libraries/io/netty/netty-handler/4.1.97.Final/netty-handler-4.1.97.Final.jar:/Users/alyx/Minecraft/Install/libraries/io/netty/netty-resolver/4.1.97.Final/netty-resolver-4.1.97.Final.jar:/Users/alyx/Minecraft/Install/libraries/io/netty/netty-transport-classes-epoll/4.1.97.Final/netty-transport-classes-epoll-4.1.97.Final.jar:/Users/alyx/Minecraft/Install/libraries/io/netty/netty-transport-native-unix-common/4.1.97.Final/netty-transport-native-unix-common-4.1.97.Final.jar:/Users/alyx/Minecraft/Install/libraries/io/netty/netty-transport/4.1.97.Final/netty-transport-4.1.97.Final.jar:/Users/alyx/Minecraft/Install/libraries/it/unimi/dsi/fastutil/8.5.12/fastutil-8.5.12.jar:/Users/alyx/Minecraft/Install/libraries/net/java/dev/jna/jna-platform/5.13.0/jna-platform-5.13.0.jar:/Users/alyx/Minecraft/Install/libraries/net/java/dev/jna/jna/5.13.0/jna-5.13.0.jar:/Users/alyx/Minecraft/Install/libraries/net/sf/jopt-simple/jopt-simple/5.0.4/jopt-simple-5.0.4.jar:/Users/alyx/Minecraft/Install/libraries/org/apache/commons/commons-compress/1.22/commons-compress-1.22.jar:/Users/alyx/Minecraft/Install/libraries/org/apache/commons/commons-lang3/3.13.0/commons-lang3-3.13.0.jar:/Users/alyx/Minecraft/Install/libraries/org/apache/httpcomponents/httpclient/4.5.13/httpclient-4.5.13.jar:/Users/alyx/Minecraft/Install/libraries/org/apache/httpcomponents/httpcore/4.4.16/httpcore-4.4.16.jar:/Users/alyx/Minecraft/Install/libraries/org/apache/logging/log4j/log4j-api/2.19.0/log4j-api-2.19.0.jar:/Users/alyx/Minecraft/Install/libraries/org/apache/logging/log4j/log4j-core/2.19.0/log4j-core-2.19.0.jar:/Users/alyx/Minecraft/Install/libraries/org/apache/logging/log4j/log4j-slf4j2-impl/2.19.0/log4j-slf4j2-impl-2.19.0.jar:/Users/alyx/Minecraft/Install/libraries/org/joml/joml/1.10.5/joml-1.10.5.jar:/Users/alyx/Minecraft/Install/libraries/org/lwjgl/lwjgl-glfw/3.3.2/lwjgl-glfw-3.3.2.jar:/Users/alyx/Minecraft/Install/libraries/org/lwjgl/lwjgl-glfw/3.3.2/lwjgl-glfw-3.3.2-natives-macos.jar:/Users/alyx/Minecraft/Install/libraries/org/lwjgl/lwjgl-glfw/3.3.2/lwjgl-glfw-3.3.2-natives-macos-arm64.jar:/Users/alyx/Minecraft/Install/libraries/org/lwjgl/lwjgl-jemalloc/3.3.2/lwjgl-jemalloc-3.3.2.jar:/Users/alyx/Minecraft/Install/libraries/org/lwjgl/lwjgl-jemalloc/3.3.2/lwjgl-jemalloc-3.3.2-natives-macos.jar:/Users/alyx/Minecraft/Install/libraries/org/lwjgl/lwjgl-jemalloc/3.3.2/lwjgl-jemalloc-3.3.2-natives-macos-arm64.jar:/Users/alyx/Minecraft/Install/libraries/org/lwjgl/lwjgl-openal/3.3.2/lwjgl-openal-3.3.2.jar:/Users/alyx/Minecraft/Install/libraries/org/lwjgl/lwjgl-openal/3.3.2/lwjgl-openal-3.3.2-natives-macos.jar:/Users/alyx/Minecraft/Install/libraries/org/lwjgl/lwjgl-openal/3.3.2/lwjgl-openal-3.3.2-natives-macos-arm64.jar:/Users/alyx/Minecraft/Install/libraries/org/lwjgl/lwjgl-opengl/3.3.2/lwjgl-opengl-3.3.2.jar:/Users/alyx/Minecraft/Install/libraries/org/lwjgl/lwjgl-opengl/3.3.2/lwjgl-opengl-3.3.2-natives-macos.jar:/Users/alyx/Minecraft/Install/libraries/org/lwjgl/lwjgl-opengl/3.3.2/lwjgl-opengl-3.3.2-natives-macos-arm64.jar:/Users/alyx/Minecraft/Install/libraries/org/lwjgl/lwjgl-stb/3.3.2/lwjgl-stb-3.3.2.jar:/Users/alyx/Minecraft/Install/libraries/org/lwjgl/lwjgl-stb/3.3.2/lwjgl-stb-3.3.2-natives-macos.jar:/Users/alyx/Minecraft/Install/libraries/org/lwjgl/lwjgl-stb/3.3.2/lwjgl-stb-3.3.2-natives-macos-arm64.jar:/Users/alyx/Minecraft/Install/libraries/org/lwjgl/lwjgl-tinyfd/3.3.2/lwjgl-tinyfd-3.3.2.jar:/Users/alyx/Minecraft/Install/libraries/org/lwjgl/lwjgl-tinyfd/3.3.2/lwjgl-tinyfd-3.3.2-natives-macos.jar:/Users/alyx/Minecraft/Install/libraries/org/lwjgl/lwjgl-tinyfd/3.3.2/lwjgl-tinyfd-3.3.2-natives-macos-arm64.jar:/Users/alyx/Minecraft/Install/libraries/org/lwjgl/lwjgl/3.3.2/lwjgl-3.3.2.jar:/Users/alyx/Minecraft/Install/libraries/org/lwjgl/lwjgl/3.3.2/lwjgl-3.3.2-natives-macos.jar:/Users/alyx/Minecraft/Install/libraries/org/lwjgl/lwjgl/3.3.2/lwjgl-3.3.2-natives-macos-arm64.jar:/Users/alyx/Minecraft/Install/libraries/org/slf4j/slf4j-api/2.0.7/slf4j-api-2.0.7.jar:/Users/alyx/Minecraft/Install/versions/fabric-loader-0.14.22-1.20.2/fabric-loader-0.14.22-1.20.2.jar",
		"-DFabricMcEmu=net.minecraft.client.main.Main",
		"-Xmx4096m",
		"-Xms256m",
		"-Dfml.ignorePatchDiscrepancies=true",
		"-Dfml.ignoreInvalidMinecraftCertificates=true",
		"-Duser.language=en",
		"-Duser.country=US",
		"-Dlog4j.configurationFile=/Users/alyx/Minecraft/Install/assets/log_configs/client-1.12.xml",
		"net.fabricmc.loader.impl.launch.knot.KnotClient",
		"--version", "fabric-loader-0.14.22-1.20.2",
		"--gameDir", "/Users/alyx/Minecraft/Instances/Boosted FPS [FABRIC] Performance Mods",
		"--assetsDir", "/Users/alyx/Minecraft/Install/assets",
		"--assetIndex", "8",
		"--accessToken", os.Getenv("ACCESS_TOKEN"),
		"--clientId", os.Getenv("CLIENT_ID"),
		"--username", os.Getenv("USERNAME"),
		"--uuid", os.Getenv("UUID"),
		"--xuid", os.Getenv("XUID"),
		"--userType", "msa",
		"--versionType", "release",
		"--width", "1920",
		"--height", "1080",
		"--quickPlayPath", "/Users/alyx/Minecraft/Install/quickPlay/java/1698398215400.json",
	)

	if err := cmd.Start(); err != nil {
		log.Fatalf("Failed to start Minecraft: %s", err)
	}
	// Wait for the game to launch
	time.Sleep(15 * time.Second)
}

func createNewWorld() {
	// Navigate the menu to create a new world
	robotgo.KeySleep = 300
	robotgo.KeyTap("down")
	robotgo.KeyTap("enter")
	robotgo.KeyTap("tab")
	robotgo.KeyTap("tab")
	robotgo.KeyTap("tab")
	robotgo.KeyTap("enter") // New game
	robotgo.TypeStr("BOT_SCREENSHOT")
	robotgo.KeyTap("tab")
	robotgo.KeyTap("tab")
	robotgo.KeyTap("tab")
	robotgo.KeyTap("space") // Allow cheats
	robotgo.KeyTap("tab")
	robotgo.KeyTap("enter") // Create new world

	// Wait for world generation to complete
	time.Sleep(20 * time.Second)
}

func setupScreenshot() {
	// Set random view
	runMinecraftChatCommand("/gamemode spectator")
	robotgo.KeyTap("f1") // Hide HUD
}

func takeRandomScreenshot() {
	pos := getRandomPlayerPos()
	// Teleport the player
	// /tp @p ~ ~+10 ~ 360 90 => Teleport player 10 blocks above + rotate player view to x y (horizontal (0 / 360) / vertical (90 / 0 / -90))
	runMinecraftChatCommand(fmt.Sprintf("/tp @p %s ~%s %s %s %s", pos.x, pos.y, pos.z, pos.rx, pos.rz))

	// Make sure we are not stuck inside walls
	preloadingTime := time.Now()
	for i := 0; i < 16; i++ { // 32 = 256/8/2 to avoid too many iterations
		runMinecraftChatCommand("/execute as @p at @s unless block ~ ~-8 ~ minecraft:air run tp @s ~ ~8 ~")
	}
	time.Sleep((30 * time.Second) - preloadingTime.Sub(time.Now())) // Wait for the chunks generation and rendering

	// Take a screenshot
	robotgo.KeyTap("f2") // Take native screenshot
	time.Sleep(1 * time.Second)
}

func quitGame() {
	// Close the game
	robotgo.KeyTap("q", "cmd")
}

func getLatestScreenshot() fs.FileInfo {
	fileInfo, err := getLastCreatedFile("/Users/alyx/Minecraft/Instances/Boosted FPS [FABRIC] Performance Mods/screenshots/")
	if err != nil {
		log.Fatal(err)
	}

	if fileInfo != nil {
		fmt.Printf("Screenshot found: %s, created at %s\n", fileInfo.Name(), fileInfo.ModTime())
	} else {
		log.Fatal("Screenshot not found.")
	}

	return fileInfo
}

// Dummy function, replace with actual implementation
func postScreenshotToSocialMedia() {
	file := getLatestScreenshot()
	fmt.Println("Posting to social media:", file.Name())
	// Add code to post the screenshot to social media here
	fmt.Println("Screenshot taken and game closed.")
}

func cleanup() {
	// TODO: delete latest created world => Name them "BOT_TODELETE_XXX" to simplify the deletion process?
	// TODO: delete screenshots files (?)
}

// getLastCreatedFile takes a directory path as an argument and returns the FileInfo
// of the most recently created file in that directory. If the directory is empty,
// or if there are no files in the directory, it returns nil.
func getLastCreatedFile(dir string) (fs.FileInfo, error) {
	// Initialize variables to store information about the newest file
	var newestFile fs.FileInfo
	var newestTime time.Time

	// Walk through the directory and its subdirectories
	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		// If an error occurs, return the error and stop walking the directory
		if err != nil {
			return err
		}
		// If the current entry is a directory, skip it
		if d.IsDir() {
			return nil
		}

		// Retrieve the FileInfo of the current file
		info, err := d.Info()
		if err != nil {
			return err
		}

		// Check if the current file is newer than the newest file found so far
		if info.ModTime().After(newestTime) {
			// If it is, update newestFile and newestTime with the current file's information
			newestFile = info
			newestTime = info.ModTime()
		}
		// Continue walking through the directory
		return nil
	})

	// If an error occurred during the directory walk, return the error
	if err != nil {
		return nil, err
	}

	// Return the FileInfo of the newest file found (or nil if no files were found)
	return newestFile, nil
}

func runMinecraftChatCommand(cmd string) {
	robotgo.KeyTap("t") // Enter chat
	robotgo.TypeStr(cmd)
	robotgo.KeyTap("enter") // Run command
}

func getRandomPlayerPos() PlayerPos {
	return PlayerPos{
		x:  fmt.Sprint(rand.Intn(2000)),
		y:  fmt.Sprint(rand.Intn(16)),
		z:  fmt.Sprint(rand.Intn(2000)),
		rx: fmt.Sprint(rand.Intn(361)),
		rz: fmt.Sprint(rand.Intn(31) - 15), // Keep between -15 and +15
	}
}
