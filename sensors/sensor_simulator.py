# Simple MQTT sensor simulator: publish a base64 image periodically
import time, base64
import paho.mqtt.client as mqtt

client = mqtt.Client()
client.connect('localhost',1883,60)
img_b64 = 'data:image/jpeg;base64,' + base64.b64encode(b'test').decode()
while True:
    payload = '{"image":"%s"}' % img_b64
    client.publish('site/site1/sensors/cam1', payload)
    print('published')
    time.sleep(10)
