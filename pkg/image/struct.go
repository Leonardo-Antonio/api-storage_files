package image

import "time"

type body struct {
	Format   string `json:"format"`
	Name     string `json:"name"`
	ImageB64 string `json:"image_b64"`
}

type infoImage struct {
	Name         string    `json:"name"`
	Src          string    `json:"src"`
	Size         int64     `json:"size"`
	Modification time.Time `json:"modification"`
}
