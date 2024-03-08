# Minecraft Screenshot Bot

![Minecraft Screenshot Bot Banner](./docs/img/banner.png)

Capture and share the beauty of Minecraft landscapes! This script launches Minecraft, takes screenshots, uses an AI to describe them, and posts them to Mastodon with an alt-text.

How does it work? French article: https://blog.alyx.pink/comment-fonctionne-mon-bot-craftviews/

## 📸 Sample Screenshots

### Example 1

 <figure>
    <img src="./docs/img/screenshot_01.png">
    <blockquote>
        <figcaption>The landscape features a lush, green forest with dense tree coverage, extending into the distance. The foliage shows varying shades of green, suggesting a mix of tree types. A calm body of water with a gentle purple hue reflects the sky, bordering the forest on one side. The sky above is painted with soft pastel colors, predominantly pink and blue, indicating either dawn or dusk. Wispy clouds are scattered across the sky, adding to the serene atmosphere of the scene. The horizon is subtly obscured by a light haze, enhancing the depth of the view.</figcaption>
        <footer>— Generated using OpenAI</cite>
    </blockquote>
</figure>

### Example 2

 <figure>
    <img src="./docs/img/screenshot_02.png">
    <blockquote>
        <figcaption>This landscape is dominated by a wintry, icy environment. Jagged ice spikes of varying heights emerge from the ground, creating a rugged and cold terrain. Between these ice formations, there are patches of water reflecting the light blue sky. The horizon showcases a flat, snow-covered plain that extends into the distance. Above, the sky is clear with a light scattering of clouds, tinged with soft shades of pink and purple, suggesting a cold but tranquil atmosphere. The overall scene is serene and devoid of vegetation or wildlife, emphasizing the stark beauty of this frozen expanse.</figcaption>
        <footer>— Generated using OpenAI</cite>
    </blockquote>
</figure>

### Example 3

 <figure>
    <img src="./docs/img/screenshot_03.png">
    <blockquote>
        <figcaption>This scene depicts a verdant river valley flanked by densely wooded hills. The trees exhibit a rich green canopy, suggesting a vibrant, healthy forest. The river itself is calm, with its surface reflecting the soft glow of the sun, which appears low in the sky, hinting at either sunrise or sunset. The light casts a warm hue over the scene, enhancing the serene and peaceful mood. Mists or haze near the horizon soften the distant landscape, adding a sense of depth and mystery to the environment. The terrain is a mix of grassy areas and exposed earth, which provides a naturalistic contrast in textures and colors.</figcaption>
        <footer>— Generated using OpenAI</cite>
    </blockquote>
</figure>

### Example 4

 <figure>
    <img src="./docs/img/screenshot_04.png">
    <blockquote>
        <figcaption>The landscape features a vast, tranquil body of water that stretches to the horizon, reflecting the warm glow of the setting or rising sun. The sky is filled with a soft, pastel-colored haze, interspersed with gentle clouds. On the left, a large, cliff-like structure with a flat top looms over the water, composed of different layers and textures, suggesting geological stratification. Lush greenery adorns the coastline and some small islands, indicating a vibrant biome near the water's edge. In the distance, faint outlines of other landforms or structures can be seen, adding a sense of scale and exploration potential to the scene. The overall atmosphere is one of calmness and natural beauty.</figcaption>
        <footer>— Generated using OpenAI</cite>
    </blockquote>
</figure>

## 🌐 See the bot in action!

Check it out on Mastodon at [3615.computer/@CraftViews](https://3615.computer/@CraftViews/).

## 🧰 Tools Required

### 🧱 Curseforge/Minecraft

Of course, to teleport the player around a new world and take the screenshot, you will need the Minecraft game.

You can tweak the bot to run any Minecraft version or modpacks. I'm using a modpack tweaked for optimization, so the world generation is faster.

You can probably swap Curseforge for anything else too, as long as you can launch the game programmatically.

### 🪣 Cloudflare R2

Screenshots are archived in an R2 bucket. As of today, without any metadata.

### 🐘 Mastodon

Screenshots are published to Mastodon, on a bot account. They are using the native way to schedule a post, so you can queue some of them, [up to a limit](https://github.com/mastodon/mastodon/blob/e8605a69d22e369e34914548338c15c053db9667/app/models/scheduled_status.rb#L16-L17).

### 🥽 OpenAI

To create alt-text for the images, and make them more accessible to visually impaired people. I'm using the "GPT4 Vision Preview" model to accurately describe the landscapes.

## 🚧 Status: Proof-of-Concept

As it stands, this bot is still mostly a basic proof-of-concept built over a few weeks.

While functional, automating the entire process without tying up your computer requires further work. Today, you need to run the script manually on a machine able to render Minecraft landscapes with a high "view distance" value, including shaders, making it GPU-intensive.

Considerations for future enhancements:

- Exploring the possibility of using cloud instances with GPUs.
- I made some tests with https://www.runpod.io/, and I believe it would be possible to automate it this way, but I did not go that far.
- Implementing a more dynamic path configuration instead of the current hard-coded paths based on my laptop setup.

**Additional Information:**

- The game was installed using Curseforge.
- The modpack used is `[FABRIC] Boosted FPS`, it includes various optimizations and shaders

## ⚙️ Configuration

To configure the bot, modify the **.env** file:

```
USERNAME=your_account_username
UUID=your_account_username_uuid
ACCESS_TOKEN=account_access_token
CLIENT_ID=
XUID=

SCREENSHOTS_DIR_PATH=/path/where/the/screenshots/are/saved/

MASTODON_SERVER=https://my.instance.example
MASTODON_CLIENT_ID=
MASTODON_CLIENT_SECRET=
MASTODON_ACCESS_TOKEN=

OPENAI_API_KEY=

R2_BUCKET_NAME=
R2_PUBLIC_DOMAIN=
R2_ACCOUNT_ID=
R2_ACCESS_KEY_ID=
R2_ACCESS_KEY_SECRET=

```

> [!IMPORTANT]
> To simplify the configuration process, launch your Minecraft game as usual. Then, capture the command line used to run the game, for instance using `ps aux`. This will provide the exact details required for the bot's setup.
