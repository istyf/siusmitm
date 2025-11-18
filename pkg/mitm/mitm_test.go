package mitm_test

import (
	"bytes"
	"testing"

	"github.com/istyf/siusmitm/pkg/mitm"
	"github.com/matryer/is"
)

func TestX(t *testing.T) {
	is := is.New(t)
	is.NoErr(mitm.Process(bytes.NewBufferString(shotData)))
}

func TestPipe(t *testing.T) {
	is := is.New(t)

	result := bytes.NewBuffer(make([]byte, 0, 512))
	defer func() { is.Equal(result.Len(), len(moreData)) }()

	mitm.Pipe(t.Context(), bytes.NewBufferString(moreData), result)
}

func TestPipeWithEvenMoreData(t *testing.T) {
	is := is.New(t)

	result := bytes.NewBuffer(make([]byte, 0, 512))
	defer func() { is.Equal(result.Len(), len(evenMoreData)) }()

	mitm.Pipe(t.Context(), bytes.NewBufferString(evenMoreData), result)
}

const shotData string = `<RcData>
  <mDst>LAPTOP-P</mDst>
  <mSrc>LAPTOP-P\SiusLane\FrameTarget_1</mSrc>
  <mCmd>Binary</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>_SHOT;1;1;0;60;1;17:37:24.00;3;16;39;74;0;0;1;-0.0089;-0.0004;900;0;0;0;1103624400;61;450;0</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-P</mDst>
  <mSrc>LAPTOP-P\SiusLane\FrameTarget_1</mSrc>
  <mCmd>Binary</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>_SHOT;1;1;0;60;1;17:47:04.13;3;16;39;98;0;0;2;-0.0020;-0.0019;900;0;0;0;1103682413;61;450;0</mPrm>
</RcData>`

const moreData string = `<RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>AnswerClientInfo</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>SiusLane</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\LonDriver\StdLonInterface</mDst>
  <mSrc />
  <mCmd>LonBinary</mCmd>
  <mSnT>Local</mSnT>
  <mSbS />
  <mPrm>63 01 55 00 00 00 00 00 01 81 01 FF FF FF FF FF FF</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>RoutingTable</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPFJvdXRpbmdUYWJsZT4NCiAgICA8RGV2aWNlSUQ+TEFQVE9QLVBMUFJRNzNNXExvbkRyaXZlcjwvRGV2aWNlSUQ+DQogICAgPEFsaWFzPkxvbkRyaXZlcjwvQWxpYXM+DQogIDwvUm91dGluZ1RhYmxlPg0KICA8Um91dGluZ1RhYmxlPg0KICAgIDxEZXZpY2VJRD5MQVBUT1AtUExQUlE3M01cTG9uRHJpdmVyXFN0ZExvbkludGVyZmFjZTwvRGV2aWNlSUQ+DQogICAgPEFsaWFzPkxvbkludGVyZmFjZTwvQWxpYXM+DQogIDwvUm91dGluZ1RhYmxlPg0KICA8Um91dGluZ1RhYmxlPg0KICAgIDxEZXZpY2VJRD5MQVBUT1AtUExQUlE3M01cTG9uRHJpdmVyXFN0ZExvbkludGVyZmFjZVxGcmFtZVRhcmdldF8xIFMxIE4xMDwvRGV2aWNlSUQ+DQogICAgPEFsaWFzPkZyYW1lVGFyZ2V0XzE8L0FsaWFzPg0KICA8L1JvdXRpbmdUYWJsZT4NCiAgPFJvdXRpbmdUYWJsZT4NCiAgICA8RGV2aWNlSUQ+TEFQVE9QLVBMUFJRNzNNXExvbkRyaXZlclxTdGRMb25JbnRlcmZhY2VcVEFQMjAyXzEgUzEgTjExPC9EZXZpY2VJRD4NCiAgICA8QWxpYXM+VEFQMjAyXzE8L0FsaWFzPg0KICA8L1JvdXRpbmdUYWJsZT4NCjwvRG9jdW1lbnRFbGVtZW50Pg==</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\LonDriver</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>DriverStatus</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LonDriver</mDst>
  <mSrc />
  <mCmd>LonCommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>ResetConnections</mPrm>
</RcData><RcData>
  <mDst>FrameTarget_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>40 5B 00</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>FrameTarget_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>40 5B 00</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>FrameTarget_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>47 D1 0A</mPrm>
</RcData><RcData>
  <mDst>FrameTarget_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 C1 02 0A</mPrm>
</RcData><RcData>
  <mDst>FrameTarget_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 BF 0A</mPrm>
</RcData><RcData>
  <mDst>FrameTarget_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 9B 0A</mPrm>
</RcData><RcData>
  <mDst>FrameTarget_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 A3 0A</mPrm>
</RcData><RcData>
  <mDst>FrameTarget_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 B2 00 03 00 0A</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\LonDriver\StdLonInterface</mDst>
  <mSrc />
  <mCmd>Binary</mCmd>
  <mSnT>Local</mSnT>
  <mSbS />
  <mPrm>22 13 61 08 05 00 00 00 00 00 00 00 00 00 00 00 6D 01 00 00 06</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>FrameTarget_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 A2 00 3D 01 C1 0A</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>FrameTarget_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 9B 0A</mPrm>
</RcData><RcData>
  <mDst>FrameTarget_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 A3 0A</mPrm>
</RcData><RcData>
  <mDst>FrameTarget_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 B2 00 03 00 0A</mPrm>
</RcData><RcData>
  <mDst>FrameTarget_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 A2 00 3D 01 C1 0A</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>FrameTarget_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>40 5B 00</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>FrameTarget_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>40 5B 00</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>FrameTarget_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>40 5B 00</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>FrameTarget_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>40 5B 00</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>FrameTarget_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>40 5B 00</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>FrameTarget_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>40 5B 00</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>FrameTarget_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>40 5B 00</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>FrameTarget_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>40 5B 00</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>FrameTarget_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>40 5B 00</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>FrameTarget_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>40 5B 00</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>FrameTarget_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>40 5B 00</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>FrameTarget_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>40 5B 00</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>FrameTarget_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>40 5B 00</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane\FrameTarget_1</mSrc>
  <mCmd>Binary</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>_SHOT;1;1;0;60;1;17:35:27.74;3;16;39;97;0;0;1;0.0018;-0.0025;900;0;0;0;1103612774;61;450;0</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>FrameTarget_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>40 5B 00</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>FrameTarget_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>40 5B 00</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane\FrameTarget_1</mSrc>
  <mCmd>Binary</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>_SHOT;1;1;0;60;1;17:37:24.00;3;16;39;74;0;0;2;-0.0089;-0.0004;900;0;0;0;1103624400;61;450;0</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>FrameTarget_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>40 5B 00</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>FrameTarget_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>40 5B 00</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane\FrameTarget_1</mSrc>
  <mCmd>Binary</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>_SHOT;1;1;0;60;1;17:39:34.25;3;16;39;80;0;0;3;-0.0051;0.0055;900;0;0;0;1103637425;61;450;0</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>FrameTarget_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>40 5B 00</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane\FrameTarget_1</mSrc>
  <mCmd>Binary</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>_SHOT;1;1;0;60;1;17:40:25.61;3;16;39;65;0;0;4;-0.0111;-0.0001;900;0;0;0;1103642561;61;450;0</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane\FrameTarget_1</mSrc>
  <mCmd>Binary</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>_SHOT;1;1;0;60;1;17:41:09.56;3;16;39;98;0;0;5;-0.0026;0.0009;900;0;0;0;1103646956;61;450;0</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>FrameTarget_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>40 5B 00</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane\FrameTarget_1</mSrc>
  <mCmd>Binary</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>_SHOT;1;1;0;60;1;17:42:08.78;3;16;39;83;0;0;6;0.0066;0.0000;900;0;0;0;1103652878;61;450;0</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>FrameTarget_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>40 5B 00</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane\FrameTarget_1</mSrc>
  <mCmd>Binary</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>_SHOT;1;1;0;60;1;17:43:11.32;3;16;39;98;0;0;7;-0.0028;0.0002;900;0;0;0;1103659132;61;450;0</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>FrameTarget_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>40 5B 00</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane\FrameTarget_1</mSrc>
  <mCmd>Binary</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>_SHOT;1;1;0;60;1;17:44:01.51;3;16;39;100;0;0;8;-0.0023;-0.0002;900;0;0;0;1103664151;61;450;0</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>FrameTarget_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>40 5B 00</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>FrameTarget_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>40 5B 00</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane\FrameTarget_1</mSrc>
  <mCmd>Binary</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>_SHOT;1;1;0;60;1;17:45:51.66;3;16;39;99;0;0;9;0.0008;0.0024;900;0;0;0;1103675166;61;450;0</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>FrameTarget_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>40 5B 00</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane\FrameTarget_1</mSrc>
  <mCmd>Binary</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>_SHOT;1;1;0;60;1;17:47:04.13;3;16;39;98;0;0;10;-0.0020;-0.0019;900;0;0;0;1103682413;61;450;0</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M</mDst>
  <mSrc>LAPTOP-PLPRQ73M\SiusLane</mSrc>
  <mCmd>GetConnectedDevicesAnswer</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>PERvY3VtZW50RWxlbWVudD4NCiAgPERldmljZUluZm8+DQogICAgPERldmljZUlEPkxBUFRPUC1QTFBSUTczTVxTaXVzTGFuZTwvRGV2aWNlSUQ+DQogICAgPEluZm8+U2l1cyBBcHBsaWNhdGlvbjwvSW5mbz4NCiAgPC9EZXZpY2VJbmZvPg0KPC9Eb2N1bWVudEVsZW1lbnQ+</mPrm>
</RcData><RcData>
  <mDst>TAP202_1</mDst>
  <mSrc />
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 03 00 02</mPrm>
</RcData>`

const evenMoreData string = `<RcData>
  <mDst />
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>RequestClientInfo</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>ConnectionAccepted</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LonDriver</mSrc>
  <mCmd>DriverStatus</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>StatusReady</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LonDriver\SNI_1</mSrc>
  <mCmd>LonBinary</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>23</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LonDriver\SNI_1\HS10 S1 N10</mSrc>
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>C0 5B 00 33 30 6C 00 00 00 00</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LonDriver\SNI_1\HS10 S1 N10</mSrc>
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>C0 5B 00 33 30 6C 00 00 00 00</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LonDriver\SNI_1\HS10 S1 N10</mSrc>
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>C7 D1 00 82 00 03 00 00 00 00 00 00 00 00 07 B7 00 33 30 6C 00 41 C6 D1 AB 00</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LonDriver\SNI_1\HS10 S1 N10</mSrc>
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>88 C1</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LonDriver\SNI_1\HS10 S1 N10</mSrc>
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>88 BF 00</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LonDriver\SNI_1\HS10 S1 N10</mSrc>
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 98 01</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LonDriver\SNI_1\HS10 S1 N10</mSrc>
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>88 9B 40 1D 6E 2D 40 02 1B 95 3F DC 72 38 00 00 00 89 00 00 01 EA 01</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LonDriver\SNI_1\HS10 S1 N10</mSrc>
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>88 A3 04 31 00 01 00 35 00 04 00 33 00 03 00 34 00 00 00</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LonDriver\SNI_1\HS10 S1 N10</mSrc>
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>88 B2 00</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LonDriver\SNI_1</mSrc>
  <mCmd>LonBinary</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>2D 07 02 67 CE D0 00</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LonDriver\SNI_1\HS10 S1 N10</mSrc>
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>88 A2</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LonDriver\SNI_1\HS10 S1 N10</mSrc>
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>08 98 01</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LonDriver\SNI_1\HS10 S1 N10</mSrc>
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>88 9B 40 1D 6E 2D 40 02 1B 95 3F DC 72 38 00 00 00 88 00 00 01 EA 01</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LonDriver\SNI_1\HS10 S1 N10</mSrc>
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>88 A3 04 31 00 01 00 35 00 04 00 33 00 03 00 34 00 00 00</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LonDriver\SNI_1\HS10 S1 N10</mSrc>
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>88 B2 00</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LonDriver\SNI_1\HS10 S1 N10</mSrc>
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>88 A2</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LonDriver\SNI_1\HS10 S1 N10</mSrc>
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>C0 5B 00 33 30 6C 00 00 00 00</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LonDriver\SNI_1\HS10 S1 N10</mSrc>
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>C0 5B 00 33 30 6C 00 00 00 00</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LonDriver\SNI_1\HS10 S1 N10</mSrc>
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>C0 5B 00 33 30 6C 00 00 00 00</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LonDriver\SNI_1\HS10 S1 N10</mSrc>
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>C0 5B 00 33 30 6C 00 00 00 00</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LonDriver\SNI_1\HS10 S1 N10</mSrc>
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>C0 5B 00 33 30 6C 00 00 00 00</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LonDriver\SNI_1\HS10 S1 N10</mSrc>
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>C0 5B 00 33 30 6C 00 00 00 00</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LonDriver\SNI_1\HS10 S1 N10</mSrc>
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>C0 5B 00 33 30 6C 00 00 00 00</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LonDriver\SNI_1\HS10 S1 N10</mSrc>
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>C0 5B 00 33 30 6C 00 00 00 00</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LonDriver\SNI_1\HS10 S1 N10</mSrc>
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>C0 5B 00 33 30 6C 00 00 00 00</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LonDriver\SNI_1\HS10 S1 N10</mSrc>
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>C0 5B 00 33 30 6C 00 00 00 00</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LonDriver\SNI_1\HS10 S1 N10</mSrc>
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>C0 5B 00 33 30 6C 00 00 00 00</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LonDriver\SNI_1\HS10 S1 N10</mSrc>
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>C0 5B 00 33 30 6C 00 00 00 00</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LonDriver\SNI_1\HS10 S1 N10</mSrc>
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>C0 5B 00 33 30 6C 00 00 00 00</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LonDriver\SNI_1\HS10 S1 N10</mSrc>
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>03 FC 00 01 41 C7 CB 66 3A E7 6D C3 BB 24 97 03 00 00 00 00 00 00 00 00 00 07 00 00 00 00 00 00 00 FF FF FF 82 00</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LonDriver\SNI_1\HS10 S1 N10</mSrc>
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>C0 5B 00 33 30 6C 00 00 00 00</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LonDriver\SNI_1\HS10 S1 N10</mSrc>
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>C0 5B 00 33 30 6C 00 00 00 00</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LonDriver\SNI_1\HS10 S1 N10</mSrc>
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>03 FC 00 01 41 C7 F8 D0 BC 11 65 02 B9 D9 CF 6B 00 00 00 00 00 00 00 00 00 07 00 00 00 00 00 00 00 FF FF FF 82 00</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LonDriver\SNI_1\HS10 S1 N10</mSrc>
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>C0 5B 00 33 30 6C 00 00 00 00</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LonDriver\SNI_1\HS10 S1 N10</mSrc>
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>C0 5B 00 33 30 6C 00 00 00 00</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LonDriver\SNI_1\HS10 S1 N10</mSrc>
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>03 FC 00 01 41 C8 2B B1 BB A5 B1 48 3B B3 E6 F0 00 00 00 00 00 00 00 00 00 07 00 00 00 00 00 00 00 FF FF FF 82 00</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LonDriver\SNI_1\HS10 S1 N10</mSrc>
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>C0 5B 00 33 30 6C 00 00 00 00</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LonDriver\SNI_1\HS10 S1 N10</mSrc>
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>03 FC 00 01 41 C8 3F C1 BC 36 A4 86 B8 D1 84 B0 00 00 00 00 00 00 00 00 00 07 00 00 00 00 00 00 00 FF FF FF 82 00</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LonDriver\SNI_1\HS10 S1 N10</mSrc>
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>03 FC 00 01 41 C8 50 EC BB 2C 57 AF 3A 69 A2 E5 00 00 00 00 00 00 00 00 00 07 00 00 00 00 00 00 00 FF FF FF 82 00</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LonDriver\SNI_1\HS10 S1 N10</mSrc>
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>C0 5B 00 33 30 6C 00 00 00 00</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LonDriver\SNI_1\HS10 S1 N10</mSrc>
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>03 FC 00 01 41 C8 68 0E 3B D8 70 33 B8 4E 91 03 00 00 00 00 00 00 00 00 00 07 00 00 00 00 00 00 00 FF FF FF 82 00</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LonDriver\SNI_1\HS10 S1 N10</mSrc>
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>C0 5B 00 33 30 6C 00 00 00 00</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LonDriver\SNI_1\HS10 S1 N10</mSrc>
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>03 FC 00 01 41 C8 80 7C BB 39 80 D8 39 81 F6 03 00 00 00 00 00 00 00 00 00 07 00 00 00 00 00 00 00 FF FF FF 82 00</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LonDriver\SNI_1\HS10 S1 N10</mSrc>
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>C0 5B 00 33 30 6C 00 00 00 00</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LonDriver\SNI_1\HS10 S1 N10</mSrc>
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>03 FC 00 01 41 C8 94 17 BB 15 6E CA B9 69 0C 65 00 00 00 00 00 00 00 00 00 07 00 00 00 00 00 00 00 FF FF FF 82 00</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LonDriver\SNI_1\HS10 S1 N10</mSrc>
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>C0 5B 00 33 30 6C 00 00 00 00</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LonDriver\SNI_1\HS10 S1 N10</mSrc>
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>C0 5B 00 33 30 6C 00 00 00 00</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LonDriver\SNI_1\HS10 S1 N10</mSrc>
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>03 FC 00 01 41 C8 BF 1E 3A 4F 3C 81 3B 1D 79 F9 00 00 00 00 00 00 00 00 00 07 00 00 00 00 00 00 00 FF FF FF 82 00</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LonDriver\SNI_1\HS10 S1 N10</mSrc>
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>C0 5B 00 33 30 6C 00 00 00 00</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LonDriver\SNI_1\HS10 S1 N10</mSrc>
  <mCmd>RCICommand</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm>03 FC 00 01 41 C8 DB 6D BB 04 90 4F BA F8 4D FC 00 00 00 00 00 00 00 00 00 07 00 00 00 00 00 00 00 FF FF FF 82 00</mPrm>
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData><RcData>
  <mDst>LAPTOP-PLPRQ73M\SiusLane</mDst>
  <mSrc>LAPTOP-PLPRQ73M</mSrc>
  <mCmd>GetConnectedDevices</mCmd>
  <mSnT>SingleNode</mSnT>
  <mSbS />
  <mPrm />
</RcData>`
