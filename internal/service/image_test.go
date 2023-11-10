package service

import (
	"github.com/stretchr/testify/require"
	"smolathon/internal/entity"
	"testing"
)

func Test_keyLinkMapper(t *testing.T) {
	testCases := []struct {
		name string
		data []keyLink
		want []entity.Image
	}{
		{
			name: "single size",
			data: []keyLink{
				{
					Key:  "dbb2edca-b8fb-4d9d-bb65-1461a6fed8c9/preview/castle-s.png",
					Link: "https://storage.lechefer.ru/dbb2edca-b8fb-4d9d-bb65-1461a6fed8c9/preview/castle-m.png",
				},
			},
			want: []entity.Image{
				{
					Sizes: entity.Sizes{
						M: entity.ImageSize{
							Size: entity.M,
							Url:  "https://storage.lechefer.ru/dbb2edca-b8fb-4d9d-bb65-1461a6fed8c9/preview/castle-m.png",
						},
					},
				},
			},
		},
		{
			name: "many size",
			data: []keyLink{
				{
					Key:  "afcfe1f3-2599-4a84-b814-f3bc12fb09d8/6c35331c-9df1-4257-8b91-1be0e355bc2c/castle-s.png",
					Link: "https://storage.lechefer.ru/afcfe1f3-2599-4a84-b814-f3bc12fb09d8/6c35331c-9df1-4257-8b91-1be0e355bc2c/castle-m.png",
				},
				{
					Key:  "afcfe1f3-2599-4a84-b814-f3bc12fb09d8/6c35331c-9df1-4257-8b91-1be0e355bc2c/castle-q.png",
					Link: "https://storage.lechefer.ru/afcfe1f3-2599-4a84-b814-f3bc12fb09d8/6c35331c-9df1-4257-8b91-1be0e355bc2c/castle-x.png",
				},
			},
			want: []entity.Image{
				{
					Sizes: entity.Sizes{
						M: entity.ImageSize{
							Size: entity.M,
							Url:  "https://storage.lechefer.ru/afcfe1f3-2599-4a84-b814-f3bc12fb09d8/6c35331c-9df1-4257-8b91-1be0e355bc2c/castle-m.png",
						},
						X: entity.ImageSize{
							Size: entity.X,
							Url:  "https://storage.lechefer.ru/afcfe1f3-2599-4a84-b814-f3bc12fb09d8/6c35331c-9df1-4257-8b91-1be0e355bc2c/castle-x.png",
						},
					},
				},
			},
		},
		{
			name: "many image",
			data: []keyLink{
				{
					Key:  "afcfe1f3-2599-4a84-b814-f3bc12fb09d8/6c35331c-9df1-4257-8b91-1be0e355bc2c/castle-s.png",
					Link: "https://storage.lechefer.ru/afcfe1f3-2599-4a84-b814-f3bc12fb09d8/6c35331c-9df1-4257-8b91-1be0e355bc2c/castle-m.png",
				},
				{
					Key:  "afcfe1f3-2599-4a84-b814-f3bc12fb09d8/d297d769-cab7-4604-93a5-64101fe0df42/forest-q.png",
					Link: "https://storage.lechefer.ru/afcfe1f3-2599-4a84-b814-f3bc12fb09d8/d297d769-cab7-4604-93a5-64101fe0df42/forest-m.png",
				},
			},
			want: []entity.Image{
				{
					Sizes: entity.Sizes{
						M: entity.ImageSize{
							Size: entity.M,
							Url:  "https://storage.lechefer.ru/afcfe1f3-2599-4a84-b814-f3bc12fb09d8/6c35331c-9df1-4257-8b91-1be0e355bc2c/castle-m.png",
						},
					},
				},
				{
					Sizes: entity.Sizes{
						X: entity.ImageSize{
							Size: entity.X,
							Url:  "https://storage.lechefer.ru/afcfe1f3-2599-4a84-b814-f3bc12fb09d8/d297d769-cab7-4604-93a5-64101fe0df42/forest-m.png",
						},
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := keyLinkMapper(tc.data)
			require.Equal(t, tc.want, actual)
		})
	}
}
