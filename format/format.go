package format

import (
	"github.com/dongjialong2006/vdk/av/avutil"
	"github.com/dongjialong2006/vdk/format/aac"
	"github.com/dongjialong2006/vdk/format/flv"
	"github.com/dongjialong2006/vdk/format/mp4"
	"github.com/dongjialong2006/vdk/format/rtmp"
	"github.com/dongjialong2006/vdk/format/rtsp"
	"github.com/dongjialong2006/vdk/format/ts"
)

func RegisterAll() {
	avutil.DefaultHandlers.Add(mp4.Handler)
	avutil.DefaultHandlers.Add(ts.Handler)
	avutil.DefaultHandlers.Add(rtmp.Handler)
	avutil.DefaultHandlers.Add(rtsp.Handler)
	avutil.DefaultHandlers.Add(flv.Handler)
	avutil.DefaultHandlers.Add(aac.Handler)
}
