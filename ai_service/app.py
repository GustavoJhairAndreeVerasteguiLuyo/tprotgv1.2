from fastapi import FastAPI
from pydantic import BaseModel
from typing import List
import base64
import cv2
import numpy as np

class InferReq(BaseModel):
    image: str

class Detection(BaseModel):
    label: str
    confidence: float

class InferResp(BaseModel):
    detections: List[Detection]

app = FastAPI()

def detect_faces_and_motion(image_b64: str):
    # Decode base64 image
    img_data = base64.b64decode(image_b64.split(',')[-1])
    nparr = np.frombuffer(img_data, np.uint8)
    img = cv2.imdecode(nparr, cv2.IMREAD_COLOR)
    # Placeholder: simple face detection with OpenCV's Haar cascades (if available)
    gray = cv2.cvtColor(img, cv2.COLOR_BGR2GRAY)
    face_cascade = cv2.CascadeClassifier(cv2.data.haarcascades + 'haarcascade_frontalface_default.xml')
    faces = face_cascade.detectMultiScale(gray, 1.1, 4)
    detections = []
    for (x,y,w,h) in faces:
        detections.append({'label':'face','confidence':0.9})
    # Motion detection placeholder: we return a dummy motion detection
    detections.append({'label':'motion','confidence':0.6})
    return detections

@app.post('/infer', response_model=InferResp)
async def infer(req: InferReq):
    dets = detect_faces_and_motion(req.image)
    return InferResp(detections=[Detection(label=d['label'], confidence=d['confidence']) for d in dets])
