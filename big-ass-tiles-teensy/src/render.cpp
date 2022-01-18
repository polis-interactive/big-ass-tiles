
#include "render.hpp"

#include "renderers/blinker.hpp"

#include <Arduino.h>


CRGB RenderController::render_leds_[COLUMNS][ROWS];
RenderControllerStates RenderController::state_ = RenderControllerStates::BLINKER;

long _timestamp;


void RenderController::SetupRenderController() {
    _timestamp = millis();
    BlinkerRenderer::SetupRenderer(_timestamp);
}


void RenderController::RunRenderController() {
    runRenderer();
    setOutput();
}

void RenderController::runRenderer() {
    _timestamp = millis();
    switch (state_)
    {
    case RenderControllerStates::BLINKER:
        BlinkerRenderer::RunRenderer(_timestamp);
        break;
    default:
        break;
    }
}

void RenderController::setOutput() {
    for (int x= 0; x < COLUMNS; x++) {
        for (int y = 0; y < ROWS; y++) {
            LedController::SetCellColor(x, y, render_leds_[x][y]);
        }
    }
}