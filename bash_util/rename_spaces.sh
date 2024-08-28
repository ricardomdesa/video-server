#!/bin/bash
for f in *.mp4; do mv "$f" "${f// /_}"; done
