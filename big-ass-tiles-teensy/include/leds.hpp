
#ifndef LEDS_H
#define LEDS_H


#include <FastLED.h>

#define DATA_PIN    1
#define LED_TYPE    WS2811
#define COLOR_ORDER RGB
#define BRIGHTNESS  191
#define CORRECTION  TypicalPixelString

// 50 for 20 hz, 33 for 30 hz
#define LED_REFRESH_MILLIS 50

// #define ROWS            3
// #define COLUMNS         11
// #define LEDS_PER_CELL   8
#define ROWS            2
#define COLUMNS         3
#define LEDS_PER_CELL   2

#define LEDS_COUNT      ROWS * COLUMNS * LEDS_PER_CELL

#define LEDS_PER_ROW    COLUMNS * LEDS_PER_CELL

#define LEDS_PER_COLUMN_SCOOT   2
#define LEDS_PER_ROW_SCOOT      LEDS_PER_COLUMN_SCOOT * COLUMNS

class LedController {
public:


    static void SetupLedController();
    static void RunLedController();

    static void SetLedsColor(const struct CRGB& color);
    static void SetCellColor(const int x, const int y, const struct CRGB& color);

private:
    static void setupFastLed();
    static void setupLedGrid();
    static long last_update_timestamp_;
    static CRGB leds_[LEDS_COUNT];
    static CRGB *led_grid_[COLUMNS][ROWS][LEDS_PER_CELL];
};


#endif /* LEDS_H */
