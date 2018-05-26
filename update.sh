rm views/index.html
rm -rf static/js
rm -rf static/css

mv front/dist/index.html views/
mv front/dist/static/js static/
mv front/dist/static/css static/