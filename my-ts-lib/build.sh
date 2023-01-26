#!/usr/bin/env sh
npm run build
rm -f bundle_dist/*
mkdir bundle_dist
cp dist/assets/*.js bundle_dist/bundle.js
