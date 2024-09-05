package media

import "github.com/l1qwie/Fmtogram/types"

func (ph *Photo) WriteCaption(caption string) {
	ph.Caption = caption
}

func (ph *Photo) WriteParseMode(parsemode string) {
	ph.ParseMode = parsemode
}

func (ph *Photo) WriteCaptionEntities(captionEntities []*types.MessageEntity) {
	ph.CaptionEntities = captionEntities
}

func (ph *Photo) WriteShowCaptionAboveMedia() {
	ph.ShowCaptionAboveMedia = true
}

func (ph *Photo) WriteHasSpoiler() {
	ph.HasSpoiler = true
}

func (vd *Video) WriteThumbnail(thumbnail string) {
	vd.Thumbnail = thumbnail
}

func (vd *Video) WriteCaption(caption string) {
	vd.Caption = caption
}

func (vd *Video) WriteParseMode(parsemode string) {
	vd.ParseMode = parsemode
}

func (vd *Video) WriteCaptionEntities(captionEntities []*types.MessageEntity) {
	vd.CaptionEntities = captionEntities
}

func (vd *Video) WriteShowCaptionAboveMedia() {
	vd.ShowCaptionAboveMedia = true
}

func (vd *Video) WriteWidth(width int) {
	vd.Width = width
}

func (vd *Video) WriteHeight(height int) {
	vd.Height = height
}

func (vd *Video) WriteDuration(duration int) {
	vd.Duration = duration
}

func (vd *Video) WriteSupportsStreaming() {
	vd.SupportsStreaming = true
}

func (vd *Video) WriteHasSpoiler() {
	vd.HasSpoiler = true
}

func (a *Audio) WriteThumbnail(thumbnail string) {
	a.Thumbnail = thumbnail
}

func (ad *Audio) WriteCaption(caption string) {
	ad.Caption = caption
}

func (ad *Audio) WriteParseMode(parsemode string) {
	ad.ParseMode = parsemode
}

func (ad *Audio) WriteCaptionEntities(captionEntities []*types.MessageEntity) {
	ad.CaptionEntities = captionEntities
}

func (ad *Audio) WriteDuration(duration int) {
	ad.Duration = duration
}

func (ad *Audio) WritePerformer(performer string) {
	ad.Performer = performer
}

func (ad *Audio) WriteTitle(title string) {
	ad.Title = title
}

func (dc *Document) WriteThumbnail(thumbnail string) {
	dc.Thumbnail = thumbnail
}

func (dc *Document) WriteCaption(caption string) {
	dc.Caption = caption
}

func (dc *Document) WriteParseMode(parsemode string) {
	dc.ParseMode = parsemode
}

func (dc *Document) WriteCaptionEntities(captionEntities []*types.MessageEntity) {
	dc.CaptionEntities = captionEntities
}

func (dc *Document) WriteDisableContentTypeDetection() {
	dc.DisableContentTypeDetection = true
}
