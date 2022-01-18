#include <Arduino.h>

#include "leds.hpp"
#include "render.hpp"


void setup() {
    LedController::SetupLedController();
    RenderController::SetupRenderController();
}

void loop() {
    RenderController::RunRenderController();
    LedController::RunLedController();
}

