package media

import (
	"github.com/l1qwie/Fmtogram/logs"
	"github.com/l1qwie/Fmtogram/types"
)

const (
	objectPhoto    string = "Photo"
	objectVideo    string = "Video"
	objectAudio    string = "Audio"
	objectDocument string = "Document"
	checkString    int    = -1
	checkArray     int    = -2
	checkBool      int    = -3
	checkInt       int    = -4
)

func (ph *Photo) WritePhoto(photo string) {
	if !isItEmply(ph, checkString, ph.Photo) {
		logs.DataIsntEmply(objectPhoto, "Photo", ph.Photo)
	}
	ph.Photo = photo
	ph.Media = photo
	logs.DataWrittenSuccessfully(objectPhoto, "Photo")
}

func (ph *Photo) WriteCaption(caption string) {
	if !isItEmply(ph, checkString, ph.Caption) {
		logs.DataIsntEmply(objectPhoto, "Caption", ph.Caption)
	}
	ph.Caption = caption
	logs.DataWrittenSuccessfully(objectPhoto, "Caption")
}

func (ph *Photo) WriteParseMode(parsemode string) {
	if !isItEmply(ph, checkString, ph.ParseMode) {
		logs.DataIsntEmply(objectPhoto, "Parse Mode", ph.ParseMode)
	}
	ph.ParseMode = parsemode
	logs.DataWrittenSuccessfully(objectPhoto, "Parse Mode")
}

func (ph *Photo) WriteCaptionEntities(captionEntities []*types.MessageEntity) {
	if !isItEmply(ph, checkArray, ph.CaptionEntities) {
		logs.DataIsntEmply(objectPhoto, "Caption Entities", ph.CaptionEntities)
	}
	ph.CaptionEntities = captionEntities
	logs.DataWrittenSuccessfully(objectPhoto, "Caption Entities")
}

func (ph *Photo) WriteShowCaptionAboveMedia() {
	if !isItEmply(ph, checkBool, ph.ShowCaptionAboveMedia) {
		logs.DataIsntEmply(objectPhoto, "Show Caption Above Media", ph.ShowCaptionAboveMedia)
	}
	ph.ShowCaptionAboveMedia = true
	logs.SettedParam("Show Caption Above Media", objectPhoto, ph.ShowCaptionAboveMedia)
}

func (ph *Photo) WriteHasSpoiler() {
	if !isItEmply(ph, checkBool, ph.HasSpoiler) {
		logs.DataIsntEmply(objectPhoto, "Has Spoiler", ph.HasSpoiler)
	}
	ph.HasSpoiler = true
	logs.SettedParam("Has Spoiler", objectPhoto, ph.HasSpoiler)
}

func (ph *Photo) WriteGottenFrom(gottenfrom int) {
	if !isItEmply(ph, checkInt, ph.gottenFrom) {
		logs.DataIsntEmply(objectPhoto, "Gotten From", ph.gottenFrom)
	}
	ph.gottenFrom = gottenfrom
	logs.DataWrittenSuccessfully(objectPhoto, "Gotten From")
}

func (vd *Video) WriteThumbnail(thumbnail string) {
	if !isItEmply(vd, checkString, vd.Thumbnail) {
		logs.DataIsntEmply("Video", "Thumbnail", vd.Thumbnail)
	}
	vd.Thumbnail = thumbnail
	logs.DataWrittenSuccessfully(objectVideo, "Thumbnail")
}

func (vd *Video) WriteCaption(caption string) {
	if !isItEmply(vd, checkString, vd.Caption) {
		logs.DataIsntEmply("Video", "Caption", vd.Thumbnail)
	}
	vd.Caption = caption
	logs.DataWrittenSuccessfully(objectVideo, "Caption")
}

func (vd *Video) WriteParseMode(parsemode string) {
	if !isItEmply(vd, checkString, vd.ParseMode) {
		logs.DataIsntEmply("Video", "Parse Mode", vd.Thumbnail)
	}
	vd.ParseMode = parsemode
	logs.DataWrittenSuccessfully(objectVideo, "Parse Mode")
}

func (vd *Video) WriteCaptionEntities(captionEntities []*types.MessageEntity) {
	if !isItEmply(vd, checkArray, vd.CaptionEntities) {
		logs.DataIsntEmply("Video", "Caption Entities", vd.Thumbnail)
	}
	vd.CaptionEntities = captionEntities
	logs.DataWrittenSuccessfully(objectVideo, "Caption Entities")
}

func (vd *Video) WriteShowCaptionAboveMedia() {
	if !isItEmply(vd, checkBool, vd.ShowCaptionAboveMedia) {
		logs.DataIsntEmply("Video", "Show Caption Above Media", vd.ShowCaptionAboveMedia)
	}
	vd.ShowCaptionAboveMedia = true
	logs.SettedParam("Show Caption Above Media", "Video", vd.ShowCaptionAboveMedia)
}

func (vd *Video) WriteWidth(width int) {
	if !isItEmply(vd, checkInt, vd.Width) {
		logs.DataIsntEmply("Video", "Width", vd.Width)
	}
	vd.Width = width
	logs.DataWrittenSuccessfully(objectVideo, "Width")
}

func (vd *Video) WriteHeight(height int) {
	if !isItEmply(vd, checkInt, vd.Height) {
		logs.DataIsntEmply("Video", "Height", vd.Height)
	}
	vd.Height = height
	logs.DataWrittenSuccessfully(objectVideo, "Height")
}

func (vd *Video) WriteDuration(duration int) {
	if !isItEmply(vd, checkInt, vd.Duration) {
		logs.DataIsntEmply("Video", "Duration", vd.Duration)
	}
	vd.Duration = duration
	logs.DataWrittenSuccessfully(objectVideo, "Duration")
}

func (vd *Video) WriteSupportsStreaming() {
	if !isItEmply(vd, checkBool, vd.SupportsStreaming) {
		logs.DataIsntEmply("Video", "Supports Streaming", vd.SupportsStreaming)
	}
	vd.SupportsStreaming = true
	logs.SettedParam("Supports Streaming", "Video", vd.SupportsStreaming)
}

func (vd *Video) WriteHasSpoiler() {
	if !isItEmply(vd, checkBool, vd.HasSpoiler) {
		logs.DataIsntEmply("Video", "Has Spoiler", vd.HasSpoiler)
	}
	vd.HasSpoiler = true
	logs.SettedParam("Has Spoiler", "Video", vd.SupportsStreaming)
}

func (vd *Video) WriteGottenFrom(gottenfrom int) {
	if !isItEmply(vd, checkInt, vd.VideoGottenFrom) {
		logs.DataIsntEmply(objectVideo, "Gotten From", vd.VideoGottenFrom)
	}
	vd.VideoGottenFrom = gottenfrom
	logs.DataWrittenSuccessfully(objectVideo, "Gotten From")
}

func (vd *Video) WriteThumbnailGottenFrom(gottenfrom int) {
	if !isItEmply(vd, checkInt, vd.ThumbnailGottenFrom) {
		logs.DataIsntEmply(objectVideo, "Thumbnail Gotten From", vd.ThumbnailGottenFrom)
	}
	vd.ThumbnailGottenFrom = gottenfrom
	logs.DataWrittenSuccessfully(objectVideo, "Thumbnail Gotten From")
}

func (ad *Audio) WriteThumbnail(thumbnail string) {
	if !isItEmply(ad, checkString, ad.Thumbnail) {
		logs.DataIsntEmply("Audio", "Thumbnail", ad.Thumbnail)
	}
	ad.Thumbnail = thumbnail
	logs.DataWrittenSuccessfully(objectAudio, "Thumbnail")
}

func (ad *Audio) WriteCaption(caption string) {
	if !isItEmply(ad, checkString, ad.Caption) {
		logs.DataIsntEmply("Audio", "Caption", ad.Caption)
	}
	ad.Caption = caption
	logs.DataWrittenSuccessfully(objectAudio, "Caption")
}

func (ad *Audio) WriteParseMode(parsemode string) {
	if !isItEmply(ad, checkString, ad.ParseMode) {
		logs.DataIsntEmply("Audio", "Parse Mode", ad.ParseMode)
	}
	ad.ParseMode = parsemode
	logs.DataWrittenSuccessfully(objectAudio, "Parse Mode")
}

func (ad *Audio) WriteCaptionEntities(captionEntities []*types.MessageEntity) {
	if !isItEmply(ad, checkArray, ad.CaptionEntities) {
		logs.DataIsntEmply("Audio", "Caption Entities", ad.CaptionEntities)
	}
	ad.CaptionEntities = captionEntities
	logs.DataWrittenSuccessfully(objectAudio, "Caption Entities")
}

func (ad *Audio) WriteDuration(duration int) {
	if !isItEmply(ad, checkInt, ad.Duration) {
		logs.DataIsntEmply("Audio", "Duration", ad.Duration)
	}
	ad.Duration = duration
	logs.DataWrittenSuccessfully(objectAudio, "Duration")
}

func (ad *Audio) WritePerformer(performer string) {
	if !isItEmply(ad, checkString, ad.Performer) {
		logs.DataIsntEmply("Audio", "Performer", ad.Performer)
	}
	ad.Performer = performer
	logs.DataWrittenSuccessfully(objectAudio, "Performer")
}

func (ad *Audio) WriteTitle(title string) {
	if !isItEmply(ad, checkString, ad.Title) {
		logs.DataIsntEmply("Audio", "Title", ad.Title)
	}
	ad.Title = title
	logs.DataWrittenSuccessfully(objectAudio, "Title")
}

func (ad *Audio) WriteGottenFrom(gottenfrom int) {
	if !isItEmply(ad, checkInt, ad.AudioGottenFrom) {
		logs.DataIsntEmply(objectAudio, "Gotten From", ad.AudioGottenFrom)
	}
	ad.AudioGottenFrom = gottenfrom
	logs.DataWrittenSuccessfully(objectAudio, "Gotten From")
}

func (ad *Audio) WriteThumbnailGottenFrom(gottenfrom int) {
	if !isItEmply(ad, checkInt, ad.ThumbnailGottenFrom) {
		logs.DataIsntEmply(objectAudio, "Thumbnail Gotten From", ad.ThumbnailGottenFrom)
	}
	ad.ThumbnailGottenFrom = gottenfrom
	logs.DataWrittenSuccessfully(objectAudio, "Thumbnail Gotten From")
}

func (dc *Document) WriteDocument(document string) {
	if !isItEmply(dc, checkString, dc.Document) {
		logs.DataIsntEmply("Document", "Document", dc.Document)
	}
	dc.Document = document
	logs.DataWrittenSuccessfully(objectDocument, "Document")
}

func (dc *Document) WriteThumbnail(thumbnail string) {
	if !isItEmply(dc, checkString, dc.Thumbnail) {
		logs.DataIsntEmply("Document", "Thumbnail", dc.Thumbnail)
	}
	dc.Thumbnail = thumbnail
	logs.DataWrittenSuccessfully(objectDocument, "Thumbnail")
}

func (dc *Document) WriteCaption(caption string) {
	if !isItEmply(dc, checkString, dc.Caption) {
		logs.DataIsntEmply("Document", "Caption", dc.Caption)
	}
	dc.Caption = caption
	logs.DataWrittenSuccessfully(objectDocument, "Caption")
}

func (dc *Document) WriteParseMode(parsemode string) {
	if !isItEmply(dc, checkString, dc.ParseMode) {
		logs.DataIsntEmply("Document", "Parse Mode", dc.ParseMode)
	}
	dc.ParseMode = parsemode
	logs.DataWrittenSuccessfully(objectDocument, "Parse Mode")
}

func (dc *Document) WriteCaptionEntities(captionEntities []*types.MessageEntity) {
	if !isItEmply(dc, checkArray, dc.CaptionEntities) {
		logs.DataIsntEmply("Document", "Caption Entities", dc.CaptionEntities)
	}
	dc.CaptionEntities = captionEntities
	logs.DataWrittenSuccessfully(objectDocument, "Caption Entities")
}

func (dc *Document) WriteDisableContentTypeDetection() {
	if !isItEmply(dc, checkArray, dc.DisableContentTypeDetection) {
		logs.DataIsntEmply("Document", "Disable Content Type Detection", dc.DisableContentTypeDetection)
	}
	dc.DisableContentTypeDetection = true
	logs.SettedParam("Disable Content Type Detection", "Document", dc.DisableContentTypeDetection)
}

func (dc *Document) WriteGottenFrom(gottenfrom int) {
	if !isItEmply(dc, checkInt, dc.DocumentGottenFrom) {
		logs.DataIsntEmply(objectDocument, "Gotten From", dc.DocumentGottenFrom)
	}
	dc.DocumentGottenFrom = gottenfrom
	logs.DataWrittenSuccessfully(objectDocument, "Gotten From")
}

func (dc *Document) WriteThumbnailGottenFrom(gottenfrom int) {
	if !isItEmply(dc, checkInt, dc.ThumbnailGottenFrom) {
		logs.DataIsntEmply(objectDocument, "Thumbnail Gotten From", dc.ThumbnailGottenFrom)
	}
	dc.ThumbnailGottenFrom = gottenfrom
	logs.DataWrittenSuccessfully(objectDocument, "Thumbnail Gotten From")
}
