package service

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"

	"github.com/steveanlorn/learning-go/testing/04_integration_test/example/internal/searchapi"
)

const succeed = "\u2713"
const failed = "\u2717"

func TestService_GetMusicByTitle(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := NewMockStore(ctrl)

	testCases := []struct {
		label          string
		title          string
		setMock        func()
		expectedResult []searchapi.Music
		expectedErr    bool
	}{
		{
			label: "success",
			title: "love",
			setMock: func() {
				mockStore.EXPECT().GetMusicByTitle(gomock.Any(), "love").
					Return([]Music{
						{
							ID:    1,
							Title: "How Deep Is Your Love",
						},
					}, nil)
			},
			expectedResult: []searchapi.Music{
				{
					ID:    1,
					Title: "How Deep Is Your Love",
				},
			},
			expectedErr: false,
		},
	}

	t.Logf("Given the need to test GetMusicByTitle")
	for _, tc := range testCases {
		tf := func(t *testing.T) {
			tc.setMock()
			service := NewService(mockStore)

			result, err := service.GetMusicByTitle(context.Background(), tc.title)
			if (err != nil) != tc.expectedErr {
				t.Fatalf("\t%s\tShould able to get music by title: %v", failed, err)
			}
			t.Logf("\t%s\tShould able to get music by title", succeed)

			if diff := cmp.Diff(result, tc.expectedResult); diff != "" {
				t.Logf("\t%s\tShould get expected music", failed)
				t.Fatal(diff)
			}
			t.Logf("\t%s\tShould get expected music", succeed)
		}
		t.Run(tc.label, tf)
	}
}
