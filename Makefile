install:
	cp build/bin/${BINNAME}_${GOOS}_${GOARCH} /usr/bin/${BINNAME}
	cp dist/debian/logo.png /usr/share/icons/hicolor/scalable/apps/dostestportal.png
	cp dist/debian/dostestportal.desktop /usr/share/applications