package gordon

import (
	"testing"
)

func common(t *testing.T, data *[]byte, expectedMethod Method) {
	expectedProjectId := "training-main-317518"
	expectedZone := "us-central1-a"
	expectedInstanceName := "instance-1"
	gotMethod, gotProjectId, gotZone, gotInstanceName, err := parse(data)
	if err != nil {
		t.Error("err != nil")
		return
	}
	if *gotMethod != expectedMethod {
		t.Errorf("gotMethod = %d; want %d", *gotMethod, expectedMethod)
	}
	if *gotProjectId != expectedProjectId {
		t.Errorf("gotProjectId = %s; want %s", *gotProjectId, expectedProjectId)
	}
	if *gotZone != expectedZone {
		t.Errorf("gotZone = %s; want %s", *gotZone, expectedZone)
	}
	if *gotInstanceName != expectedInstanceName {
		t.Errorf("goInstancetName = %s; want %s", *gotInstanceName, expectedInstanceName)
	}
}

func TestParseInsert(t *testing.T) {
	data := []byte(`{"message":{"attributes":{"logging.googleapis.com/timestamp":"2021-07-27T11:01:08.639002Z"},"data":"eyJpbnNlcnRJZCI6InNhdm1tNmNoMXMiLCJsb2dOYW1lIjoicHJvamVjdHMvdHJhaW5pbmctbWFpbi0zMTc1MTgvbG9ncy9jbG91ZGF1ZGl0Lmdvb2dsZWFwaXMuY29tJTJGYWN0aXZpdHkiLCJvcGVyYXRpb24iOnsiaWQiOiJvcGVyYXRpb24tMTYyNzM4MzY1ODUwNS01YzgxOGM4ZTBlM2Q5LWEyYmY5OTgxLTU2ZjY3MzcxIiwibGFzdCI6dHJ1ZSwicHJvZHVjZXIiOiJjb21wdXRlLmdvb2dsZWFwaXMuY29tIn0sInByb3RvUGF5bG9hZCI6eyJAdHlwZSI6InR5cGUuZ29vZ2xlYXBpcy5jb20vZ29vZ2xlLmNsb3VkLmF1ZGl0LkF1ZGl0TG9nIiwiYXV0aGVudGljYXRpb25JbmZvIjp7InByaW5jaXBhbEVtYWlsIjoidHJhaW5pbmcyLmxhcmtpbnR1Y2tlcmxsY0BnbWFpbC5jb20ifSwibWV0aG9kTmFtZSI6ImJldGEuY29tcHV0ZS5pbnN0YW5jZXMuaW5zZXJ0IiwicmVxdWVzdCI6eyJAdHlwZSI6InR5cGUuZ29vZ2xlYXBpcy5jb20vY29tcHV0ZS5pbnN0YW5jZXMuaW5zZXJ0In0sInJlcXVlc3RNZXRhZGF0YSI6eyJjYWxsZXJJcCI6IjI0LjI1MC4xNjEuNzQiLCJjYWxsZXJTdXBwbGllZFVzZXJBZ2VudCI6Ik1vemlsbGEvNS4wIChNYWNpbnRvc2g7IEludGVsIE1hYyBPUyBYIDEwXzE1XzcpIEFwcGxlV2ViS2l0LzUzNy4zNiAoS0hUTUwsIGxpa2UgR2Vja28pIENocm9tZS85MS4wLjQ0NzIuMTY0IFNhZmFyaS81MzcuMzYsZ3ppcChnZmUpLGd6aXAoZ2ZlKSJ9LCJyZXNvdXJjZU5hbWUiOiJwcm9qZWN0cy90cmFpbmluZy1tYWluLTMxNzUxOC96b25lcy91cy1jZW50cmFsMS1hL2luc3RhbmNlcy9pbnN0YW5jZS0xIiwic2VydmljZU5hbWUiOiJjb21wdXRlLmdvb2dsZWFwaXMuY29tIn0sInJlY2VpdmVUaW1lc3RhbXAiOiIyMDIxLTA3LTI3VDExOjAxOjA5LjQxNDc0NDYxOFoiLCJyZXNvdXJjZSI6eyJsYWJlbHMiOnsiaW5zdGFuY2VfaWQiOiI2MzcwNTE4Mjk4MDQ4NzQ3OTA5IiwicHJvamVjdF9pZCI6InRyYWluaW5nLW1haW4tMzE3NTE4Iiwiem9uZSI6InVzLWNlbnRyYWwxLWEifSwidHlwZSI6ImdjZV9pbnN0YW5jZSJ9LCJzZXZlcml0eSI6Ik5PVElDRSIsInRpbWVzdGFtcCI6IjIwMjEtMDctMjdUMTE6MDE6MDguNjM5MDAyWiJ9","messageId":"2779447747553487","message_id":"2779447747553487","publishTime":"2021-07-27T11:01:10.845Z","publish_time":"2021-07-27T11:01:10.845Z"},"subscription":"projects/training-main-317518/subscriptions/gordon"}`)
	expectedMethod := Allocate
	common(t, &data, expectedMethod)
}

func TestParseStop(t *testing.T) {
	data := []byte(`{"message":{"attributes":{"logging.googleapis.com/timestamp":"2021-07-27T11:11:40.497619Z"},"data":"eyJpbnNlcnRJZCI6Ii1hajMzd3FjZTFrIiwibG9nTmFtZSI6InByb2plY3RzL3RyYWluaW5nLW1haW4tMzE3NTE4L2xvZ3MvY2xvdWRhdWRpdC5nb29nbGVhcGlzLmNvbSUyRmFjdGl2aXR5Iiwib3BlcmF0aW9uIjp7ImlkIjoib3BlcmF0aW9uLTE2MjczODQyODcwNDYtNWM4MThlZTU3YTlkMy05MDdjYTE2NS03YWYyMmUzMiIsImxhc3QiOnRydWUsInByb2R1Y2VyIjoiY29tcHV0ZS5nb29nbGVhcGlzLmNvbSJ9LCJwcm90b1BheWxvYWQiOnsiQHR5cGUiOiJ0eXBlLmdvb2dsZWFwaXMuY29tL2dvb2dsZS5jbG91ZC5hdWRpdC5BdWRpdExvZyIsImF1dGhlbnRpY2F0aW9uSW5mbyI6eyJwcmluY2lwYWxFbWFpbCI6InRyYWluaW5nMi5sYXJraW50dWNrZXJsbGNAZ21haWwuY29tIn0sIm1ldGhvZE5hbWUiOiJ2MS5jb21wdXRlLmluc3RhbmNlcy5zdG9wIiwicmVxdWVzdCI6eyJAdHlwZSI6InR5cGUuZ29vZ2xlYXBpcy5jb20vY29tcHV0ZS5pbnN0YW5jZXMuc3RvcCJ9LCJyZXF1ZXN0TWV0YWRhdGEiOnsiY2FsbGVySXAiOiIyNC4yNTAuMTYxLjc0IiwiY2FsbGVyU3VwcGxpZWRVc2VyQWdlbnQiOiJNb3ppbGxhLzUuMCAoTWFjaW50b3NoOyBJbnRlbCBNYWMgT1MgWCAxMF8xNV83KSBBcHBsZVdlYktpdC81MzcuMzYgKEtIVE1MLCBsaWtlIEdlY2tvKSBDaHJvbWUvOTEuMC40NDcyLjE2NCBTYWZhcmkvNTM3LjM2LGd6aXAoZ2ZlKSxnemlwKGdmZSkifSwicmVzb3VyY2VOYW1lIjoicHJvamVjdHMvdHJhaW5pbmctbWFpbi0zMTc1MTgvem9uZXMvdXMtY2VudHJhbDEtYS9pbnN0YW5jZXMvaW5zdGFuY2UtMSIsInNlcnZpY2VOYW1lIjoiY29tcHV0ZS5nb29nbGVhcGlzLmNvbSJ9LCJyZWNlaXZlVGltZXN0YW1wIjoiMjAyMS0wNy0yN1QxMToxMTo0MC45MDY5NDYwNjRaIiwicmVzb3VyY2UiOnsibGFiZWxzIjp7Imluc3RhbmNlX2lkIjoiNjM3MDUxODI5ODA0ODc0NzkwOSIsInByb2plY3RfaWQiOiJ0cmFpbmluZy1tYWluLTMxNzUxOCIsInpvbmUiOiJ1cy1jZW50cmFsMS1hIn0sInR5cGUiOiJnY2VfaW5zdGFuY2UifSwic2V2ZXJpdHkiOiJOT1RJQ0UiLCJ0aW1lc3RhbXAiOiIyMDIxLTA3LTI3VDExOjExOjQwLjQ5NzYxOVoifQ==","messageId":"2779488393696838","message_id":"2779488393696838","publishTime":"2021-07-27T11:11:42.441Z","publish_time":"2021-07-27T11:11:42.441Z"},"subscription":"projects/training-main-317518/subscriptions/gordon"}`)
	expectedMethod := Deallocate
	common(t, &data, expectedMethod)
}

func TestParseStart(t *testing.T) {
	data := []byte(`{"message":{"attributes":{"logging.googleapis.com/timestamp":"2021-07-27T11:12:47.895356Z"},"data":"eyJpbnNlcnRJZCI6Inc3NGx2NWNoM3EiLCJsb2dOYW1lIjoicHJvamVjdHMvdHJhaW5pbmctbWFpbi0zMTc1MTgvbG9ncy9jbG91ZGF1ZGl0Lmdvb2dsZWFwaXMuY29tJTJGYWN0aXZpdHkiLCJvcGVyYXRpb24iOnsiaWQiOiJvcGVyYXRpb24tMTYyNzM4NDM1Njg0MS01YzgxOGYyODBhNjM2LWE4YjhkMjM5LThkMzNhNmUyIiwibGFzdCI6dHJ1ZSwicHJvZHVjZXIiOiJjb21wdXRlLmdvb2dsZWFwaXMuY29tIn0sInByb3RvUGF5bG9hZCI6eyJAdHlwZSI6InR5cGUuZ29vZ2xlYXBpcy5jb20vZ29vZ2xlLmNsb3VkLmF1ZGl0LkF1ZGl0TG9nIiwiYXV0aGVudGljYXRpb25JbmZvIjp7InByaW5jaXBhbEVtYWlsIjoidHJhaW5pbmcyLmxhcmtpbnR1Y2tlcmxsY0BnbWFpbC5jb20ifSwibWV0aG9kTmFtZSI6InYxLmNvbXB1dGUuaW5zdGFuY2VzLnN0YXJ0IiwicmVxdWVzdCI6eyJAdHlwZSI6InR5cGUuZ29vZ2xlYXBpcy5jb20vY29tcHV0ZS5pbnN0YW5jZXMuc3RhcnQifSwicmVxdWVzdE1ldGFkYXRhIjp7ImNhbGxlcklwIjoiMjQuMjUwLjE2MS43NCIsImNhbGxlclN1cHBsaWVkVXNlckFnZW50IjoiTW96aWxsYS81LjAgKE1hY2ludG9zaDsgSW50ZWwgTWFjIE9TIFggMTBfMTVfNykgQXBwbGVXZWJLaXQvNTM3LjM2IChLSFRNTCwgbGlrZSBHZWNrbykgQ2hyb21lLzkxLjAuNDQ3Mi4xNjQgU2FmYXJpLzUzNy4zNixnemlwKGdmZSksZ3ppcChnZmUpIn0sInJlc291cmNlTmFtZSI6InByb2plY3RzL3RyYWluaW5nLW1haW4tMzE3NTE4L3pvbmVzL3VzLWNlbnRyYWwxLWEvaW5zdGFuY2VzL2luc3RhbmNlLTEiLCJzZXJ2aWNlTmFtZSI6ImNvbXB1dGUuZ29vZ2xlYXBpcy5jb20ifSwicmVjZWl2ZVRpbWVzdGFtcCI6IjIwMjEtMDctMjdUMTE6MTI6NDguMzI1OTI3MjM1WiIsInJlc291cmNlIjp7ImxhYmVscyI6eyJpbnN0YW5jZV9pZCI6IjYzNzA1MTgyOTgwNDg3NDc5MDkiLCJwcm9qZWN0X2lkIjoidHJhaW5pbmctbWFpbi0zMTc1MTgiLCJ6b25lIjoidXMtY2VudHJhbDEtYSJ9LCJ0eXBlIjoiZ2NlX2luc3RhbmNlIn0sInNldmVyaXR5IjoiTk9USUNFIiwidGltZXN0YW1wIjoiMjAyMS0wNy0yN1QxMToxMjo0Ny44OTUzNTZaIn0=","messageId":"2779511565563134","message_id":"2779511565563134","publishTime":"2021-07-27T11:12:49.901Z","publish_time":"2021-07-27T11:12:49.901Z"},"subscription":"projects/training-main-317518/subscriptions/gordon"}`)
	expectedMethod := Allocate
	common(t, &data, expectedMethod)
}

func TestParseDelete(t *testing.T) {
	data := []byte(`{"message":{"attributes":{"logging.googleapis.com/timestamp":"2021-07-27T11:14:49.034443Z"},"data":"eyJpbnNlcnRJZCI6Ii0zMWhzb21jZXY2IiwibG9nTmFtZSI6InByb2plY3RzL3RyYWluaW5nLW1haW4tMzE3NTE4L2xvZ3MvY2xvdWRhdWRpdC5nb29nbGVhcGlzLmNvbSUyRmFjdGl2aXR5Iiwib3BlcmF0aW9uIjp7ImlkIjoib3BlcmF0aW9uLTE2MjczODQ0NzIwNDAtNWM4MThmOTVlNzM2ZS1kZTU4ZGY5YS1hM2ZhOTliYSIsImxhc3QiOnRydWUsInByb2R1Y2VyIjoiY29tcHV0ZS5nb29nbGVhcGlzLmNvbSJ9LCJwcm90b1BheWxvYWQiOnsiQHR5cGUiOiJ0eXBlLmdvb2dsZWFwaXMuY29tL2dvb2dsZS5jbG91ZC5hdWRpdC5BdWRpdExvZyIsImF1dGhlbnRpY2F0aW9uSW5mbyI6eyJwcmluY2lwYWxFbWFpbCI6InRyYWluaW5nMi5sYXJraW50dWNrZXJsbGNAZ21haWwuY29tIn0sIm1ldGhvZE5hbWUiOiJ2MS5jb21wdXRlLmluc3RhbmNlcy5kZWxldGUiLCJyZXF1ZXN0Ijp7IkB0eXBlIjoidHlwZS5nb29nbGVhcGlzLmNvbS9jb21wdXRlLmluc3RhbmNlcy5kZWxldGUifSwicmVxdWVzdE1ldGFkYXRhIjp7ImNhbGxlcklwIjoiMjQuMjUwLjE2MS43NCIsImNhbGxlclN1cHBsaWVkVXNlckFnZW50IjoiTW96aWxsYS81LjAgKE1hY2ludG9zaDsgSW50ZWwgTWFjIE9TIFggMTBfMTVfNykgQXBwbGVXZWJLaXQvNTM3LjM2IChLSFRNTCwgbGlrZSBHZWNrbykgQ2hyb21lLzkxLjAuNDQ3Mi4xNjQgU2FmYXJpLzUzNy4zNixnemlwKGdmZSksZ3ppcChnZmUpIn0sInJlc291cmNlTmFtZSI6InByb2plY3RzL3RyYWluaW5nLW1haW4tMzE3NTE4L3pvbmVzL3VzLWNlbnRyYWwxLWEvaW5zdGFuY2VzL2luc3RhbmNlLTEiLCJzZXJ2aWNlTmFtZSI6ImNvbXB1dGUuZ29vZ2xlYXBpcy5jb20ifSwicmVjZWl2ZVRpbWVzdGFtcCI6IjIwMjEtMDctMjdUMTE6MTQ6NDkuNDAyNzk1NTk2WiIsInJlc291cmNlIjp7ImxhYmVscyI6eyJpbnN0YW5jZV9pZCI6IjYzNzA1MTgyOTgwNDg3NDc5MDkiLCJwcm9qZWN0X2lkIjoidHJhaW5pbmctbWFpbi0zMTc1MTgiLCJ6b25lIjoidXMtY2VudHJhbDEtYSJ9LCJ0eXBlIjoiZ2NlX2luc3RhbmNlIn0sInNldmVyaXR5IjoiTk9USUNFIiwidGltZXN0YW1wIjoiMjAyMS0wNy0yN1QxMToxNDo0OS4wMzQ0NDNaIn0=","messageId":"2779489313138389","message_id":"2779489313138389","publishTime":"2021-07-27T11:14:49.718Z","publish_time":"2021-07-27T11:14:49.718Z"},"subscription":"projects/training-main-317518/subscriptions/gordon"}`)
	expectedMethod := Deallocate
	common(t, &data, expectedMethod)
}
