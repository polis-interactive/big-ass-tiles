
#include "renderers/blinker.hpp"
#include "render.hpp"

Blinker BlinkerRenderer::blinkers_[COLUMNS][ROWS];

void BlinkerRenderer::RunRenderer(const long timestamp) {
    for (int x = 0; x < COLUMNS; x++) {
        for (int y = 0; y < ROWS; y++) {
            RenderController::render_leds_[x][y] = blinkers_[x][y].RunBlinker(timestamp);
        }
    }
}