package service

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"net/url"
	"regexp"
	"smolathon/config"
	"smolathon/internal/entity"
	"time"
)

const _defaultExp = 24 * time.Hour

type ImageService struct {
	s3client *minio.Client

	bucket config.Bucket
}

func NewService(s3client *minio.Client, bucket config.Bucket) *ImageService {
	return &ImageService{
		s3client: s3client,
		bucket:   bucket,
	}
}

func (s *ImageService) GetQuestPreview(ctx context.Context, questId uuid.UUID) (entity.Image, error) {
	key := fmt.Sprintf("%s/preview", questId)
	keyLinks, err := s.getPresignedURLsForPost(ctx, key, _defaultExp)
	if err != nil {
		return entity.Image{}, err
	}

	preview := keyLinkMapper(keyLinks)
	if len(preview) == 0 {
		return entity.Image{}, nil
	}

	return preview[0], nil
}

func (s *ImageService) GetQuestStepImages(ctx context.Context, questId uuid.UUID, questStepId uuid.UUID) ([]entity.Image, error) {
	key := fmt.Sprintf("%s/%s", questId, questStepId)
	keyLinks, err := s.getPresignedURLsForPost(ctx, key, _defaultExp)
	if err != nil {
		return nil, err
	}

	return keyLinkMapper(keyLinks), nil
}

type keyLink struct {
	Key  string
	Link string
}

var _keyMatcherRegexp = regexp.MustCompile(`(?P<cutKey>.*)-(?P<size>.*)(?P<ext>\..*)$`)

func (s *ImageService) getPresignedURLsForPost(ctx context.Context, key string, expiration time.Duration) ([]keyLink, error) {
	opts := minio.ListObjectsOptions{
		Prefix:    key + "/",
		Recursive: true,
	}

	getCtx, cancel := context.WithCancel(ctx)
	defer cancel()
	objectInfoCh := s.s3client.ListObjects(getCtx, s.bucket.Name, opts)

	var keyLinks []keyLink
	for objectInfo := range objectInfoCh {
		if objectInfo.Err != nil {
			return nil, objectInfo.Err
		}

		objectKey := objectInfo.Key
		reqParams := make(url.Values)
		presignedURL, err := s.s3client.PresignedGetObject(getCtx, s.bucket.Name, objectKey, expiration, reqParams)
		if err != nil {
			return nil, err
		}

		keyLinks = append(keyLinks, keyLink{
			Key:  objectKey,
			Link: presignedURL.String(),
		})
	}

	return keyLinks, nil
}

func keyLinkMapper(keyLinks []keyLink) []entity.Image {
	imagesMap := make(map[string]entity.Image)
	for _, link := range keyLinks {
		matches := _keyMatcherRegexp.FindStringSubmatch(link.Key)
		sizeIdx := _keyMatcherRegexp.SubexpIndex("size")
		cutKeyIdx := _keyMatcherRegexp.SubexpIndex("cutKey")

		if len(matches) == 0 {
			continue
		}

		size := matches[sizeIdx]
		cutKey := matches[cutKeyIdx]

		switch entity.Size(size) {
		case entity.M:
			img, ok := imagesMap[cutKey]
			if !ok {
				img = entity.Image{
					Sizes: entity.Sizes{
						M: entity.ImageSize{
							Size: entity.M,
							Url:  link.Link,
						},
					},
				}
			}

			img.Sizes.M = entity.ImageSize{
				Size: entity.M,
				Url:  link.Link,
			}

			imagesMap[cutKey] = img
		case entity.X:
			img, ok := imagesMap[cutKey]
			if !ok {
				img = entity.Image{
					Sizes: entity.Sizes{
						X: entity.ImageSize{
							Size: entity.X,
							Url:  link.Link,
						},
					},
				}
			}

			img.Sizes.X = entity.ImageSize{
				Size: entity.X,
				Url:  link.Link,
			}

			imagesMap[cutKey] = img
		case entity.O:
			img, ok := imagesMap[cutKey]
			if !ok {
				img = entity.Image{
					Sizes: entity.Sizes{
						O: entity.ImageSize{
							Size: entity.O,
							Url:  link.Link,
						},
					},
				}
			}

			img.Sizes.O = entity.ImageSize{
				Size: entity.O,
				Url:  link.Link,
			}

			imagesMap[cutKey] = img
		default:
			continue
		}
	}

	var images = make([]entity.Image, 0, len(imagesMap))
	for _, image := range imagesMap {
		images = append(images, image)
	}

	return images
}
