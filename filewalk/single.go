package filewalk

// SingleFileCheck calculates some signatures on a file
func (h *Handle) SingleFileCheck(fh Handler) {
	file := fh.LoadFile(h.file)
	h.etag = fh.getEtag(file)
	h.md5sum = fh.getMd5(file)
	h.sha1sum = fh.getSha1sum(file)
}
