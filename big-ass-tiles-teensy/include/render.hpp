
#ifndef RENDER_H
#define RENDER_H


#include "leds.hpp"
#include "utility.hpp"

enum class RenderControllerStates {
    BLINKER
};

class RenderController {
public:
    static void SetupRenderController();
    static void RunRenderController();
    static CRGB render_leds_[COLUMNS][ROWS];
private:
    static void runRenderer();
    static void setOutput();
    static RenderControllerStates state_;
};


#endif /* RENDER_H */
