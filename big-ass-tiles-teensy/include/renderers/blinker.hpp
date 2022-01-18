
#ifndef BLINKER_H
#define BLINKER_H

#include "leds.hpp"

enum class BlinkerState {
    STATIC,
    FADING
};


struct Blinker {
public:

    void ResetBlinker(const long timestamp) {
        last_timestamp_ = timestamp;
        state_ = random(2) ? BlinkerState::FADING : BlinkerState::STATIC;
        if (state_ == BlinkerState::FADING) {
            hold_time_ = random(500, 5000);
        } else {
            hold_time_ = random(2000, 5000);
        }
        current_hue_ = getRandomHue();
        next_hue_ = getRandomHue();
    }

    CRGB RunBlinker(const long timestamp) {
        const long elapsed = (float) timestamp - last_timestamp_;
        if (elapsed > hold_time_) {
            doTransition(timestamp);
            return CHSV(current_hue_, 255, 255);
        }
        if (state_ == BlinkerState::STATIC) {
            return CHSV(current_hue_, 255, 255);
        } else {
            uint8_t frac = elapsed / ((float) hold_time_) * 255;
            return CHSV(lerp8by8(current_hue_, next_hue_, frac), 255, 255);
        }
    }

private:

    void doTransition(const long timestamp) {
        last_timestamp_ = timestamp;
        if (state_ == BlinkerState::FADING) {
            state_ = BlinkerState::STATIC;
            current_hue_ = next_hue_;
            hold_time_ = random(2000, 5000);
        } else {
            state_ = BlinkerState::FADING;
            next_hue_ = getRandomHue();
            hold_time_ = random(500, 5000);
        }
    }
    

    long last_timestamp_;
    long hold_time_;
    BlinkerState state_;
    uint8_t current_hue_;
    uint8_t next_hue_;

    uint8_t getRandomHue() {
        const uint8_t randomInt = random(12);
        return (randomInt / 12.0f) * 255;
    }
};

class BlinkerRenderer {
public:
    static void ResetRenderer(const long timestamp) {
        for (int x = 0; x < COLUMNS; x++) {
            for (int y = 0; y < ROWS; y++) {
                blinkers_[x][y].ResetBlinker(timestamp);
            }
        }
    }
    static void SetupRenderer(const long timestamp) {
        ResetRenderer(timestamp);
    }

    static void RunRenderer(const long timestamp);
private:
    static Blinker blinkers_[COLUMNS][ROWS];
};

#endif /* BLINKER */
