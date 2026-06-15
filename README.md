# ebitenplus

Ebitenplus is a set of tools and systems that help make ebitengine game development a little easier.

Ebitengine is (in my opinion) not really a 'game engine', it's more of a framework, it handles input, audio and GPU rendering but doesn't really provide basic higher level features like a 2D camera, centralized asset loading, UI components, scene architecture, etc.

This is where ebitenplus comes in.

## Features:

- Better asset loading (cleaner API, AssetLoader struct)
- 2D Camera (camera that can translate game world sprites relative to it's position)
- Basic scene architecture
- Vec and Rect structs (like the STL image.Point and image.Rectangle but adapted with float64 coordinates)
- Better API for text rendering with the Printer struct
- Basic ui components

**(Planned)**

## Project walkthrough:

### Directories:

- *examples* (examples for how you should use the library)
- *internal* (for now just test assets)

## Who is this for?

Anyone who wants to make games with ebitengine and doesn't want to make these basic systems from scratch everytime (mostly myself, lol).

## TODO:

- Add better defaults to the asset loader so that it doesn't just return nil and crash the game
- Make a window that can use a sprite template for it's edges
- Make more ui components (button, radio, slider) 