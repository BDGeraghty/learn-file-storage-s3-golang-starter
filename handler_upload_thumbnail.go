package main
 
import (
	"fmt"
	"net/http"
	"io"

	"github.com/bootdotdev/learn-file-storage-s3-golang-starter/internal/auth"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerUploadThumbnail(w http.ResponseWriter, r *http.Request) {
	const maxMemory int64 = 10 * 1024 * 1024 // 10 MB
	if err := r.ParseMultipartForm(maxMemory); err != nil {
		respondWithError(w, http.StatusBadRequest, "Couldn't parse form", err)
		return
	}
	file, header, err := r.FormFile("thumbnail")
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Couldn't get thumbnail file", err)
		return
	}
	defer file.Close()
	mediaType := header.Header.Get("Content-Type")
	if mediaType == "" {
		respondWithError(w, http.StatusBadRequest, "Content-Type header is missing", nil)
		return
	}	
	iodata, err := io.ReadAll(file)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't read thumbnail file", err)
		return
	}
	// GetVideo(w, r) // Removed undefined function call
	videoIDString := r.PathValue("videoID")
	videoID, err := uuid.Parse(videoIDString)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid ID", err)
		return
	}

	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't find JWT", err)
		return
	}

	userID, err := auth.ValidateJWT(token, cfg.jwtSecret)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't validate JWT", err)
		return
	}


	fmt.Println("uploading thumbnail for video", videoID, "by user", userID)

	// TODO: implement the upload here
	videoThumbnails[videoID] = thumbnail{
		data:      iodata,
		mediaType: mediaType,
	}
	respondWithJSON(w, http.StatusOK, struct {
		Message string `json:"message"`
	}{
		Message: "Thumbnail uploaded successfully",
	})
}
