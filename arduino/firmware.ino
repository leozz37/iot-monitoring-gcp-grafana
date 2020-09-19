#include "esp32-mqtt.h"

#ifdef __cplusplus
extern "C" {
#endif
  uint8_t temprature_sens_read();
#ifdef __cplusplus
}
#endif

uint8_t temprature_sens_read();

void setup() {
  // put your setup code here, to run once:
  Serial.begin(115200);
  pinMode(LED_BUILTIN, OUTPUT);
  setupCloudIoT();
}

unsigned long lastMillis = 0;
void loop() {
  mqtt->loop();
  delay(10);  // <- fixes some issues with WiFi stability

  if (!mqttClient->connected()) {
    connect();
  }

  
  if (millis() - lastMillis > 5000) {
    lastMillis = millis();

    const float temperature = (temprature_sens_read() - 32) / 1.8;

    String payload =
      String("{\"temperature\":") + String(temperature) +
      String("}");
    publishTelemetry(payload);
    
  }
}