
#include "leds.hpp"

#include <algorithm>

CRGB LedController::leds_[LEDS_COUNT];
CRGB *LedController::led_grid_[COLUMNS][ROWS][LEDS_PER_CELL] = {{{NULL}}};
long LedController::last_update_timestamp_ = 0;

void LedController::SetLedsColor(const struct CRGB &color) {
    fill_solid(leds_, LEDS_COUNT, color);
}

void LedController::SetCellColor(const int x, const int y, const struct CRGB& color) {
    for (int led = 0; led < LEDS_PER_CELL; led++) {
        *led_grid_[x][y][led] = color;
    }
}

void LedController::setupFastLed() {
    FastLED.addLeds<LED_TYPE,DATA_PIN,COLOR_ORDER>(leds_, LEDS_COUNT)
        .setCorrection(CORRECTION)
    ;

    FastLED.setBrightness(BRIGHTNESS);
}

void LedController::setupLedGrid() {
    const float leds_per_row_scoot = (float)LEDS_PER_ROW_SCOOT;
    const float leds_per_column_scoot = (float)LEDS_PER_COLUMN_SCOOT;
    for (float led_number = 0.0f; led_number < LEDS_COUNT; led_number++) {
        // casting hell
        const int row = (int)(led_number / ((float) LEDS_PER_ROW));
        const bool is_odd_row_scoot = ((int) floor(led_number / leds_per_row_scoot)) % 2; // two gives true of false, nothing to do with
                                                                             // physical parameters
        const int nominal_column = ((int) floor(led_number / leds_per_column_scoot)) % COLUMNS;
        const int actual_column = is_odd_row_scoot ? COLUMNS - nominal_column - 1 : nominal_column;
        for (int i = 0; i < LEDS_PER_CELL; i++) {
            if (led_grid_[actual_column][row][i] == NULL) {
                led_grid_[actual_column][row][i] = &leds_[(int)led_number];
                break; 
            }
        } 
    }
}

void LedController::SetupLedController() {
    setupFastLed();
    setupLedGrid();
    SetLedsColor(CRGB::Black);
    last_update_timestamp_ = millis();
}

void LedController::RunLedController() {
    const long timestamp = millis();
    const long elapsed = timestamp - last_update_timestamp_;
    if (elapsed < LED_REFRESH_MILLIS) {
        const long timeout = LED_REFRESH_MILLIS - elapsed;
        FastLED.delay(timeout);
    }
    FastLED.show();
    last_update_timestamp_ = timestamp;
}
