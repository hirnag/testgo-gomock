package mock

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/moneyforward/cayenne/app/testsupport"
)

var data = InitData()

func TestDataSet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mock := NewMockDataSet(ctrl)
	data = mock

	mock.EXPECT().Add(gomock.Any()).Return(InitData())

	exp2 := &Data{2, "jiro", "america"}
	mock.EXPECT().Find(uint8(2), "jiro").Return(exp2).AnyTimes()

	exp1 := &Data{1, "taro", "hawaii"}
	mock.EXPECT().Find(uint8(1), "taro").Return(exp1).AnyTimes()

	exp6 := &Data{6, "musashi", "ganryujima"}
	callExp6 := mock.EXPECT().Find(gomock.Any(), gomock.Any()).Return(exp6).Times(1)

	exp7 := &Data{7, "nana", "brazil"}
	mock.EXPECT().Find(gomock.Any(), gomock.Any()).Return(exp7).AnyTimes().After(callExp6)

	testsupport.AssertRow(t, "find1", data.Find(1, "taro"), exp1)
	testsupport.AssertRow(t, "find6", data.Find(3, "saburo"), exp6)
	testsupport.AssertRow(t, "find7", data.Find(9, "kuro"), exp7)
	testsupport.AssertRow(t, "find2", data.Find(2, "jiro"), exp2)
	testsupport.AssertRow(t, "find2", data.Find(2, "jiro"), exp2)
	testsupport.AssertRow(t, "find7", data.Find(3, "saburo"), exp7)
	testsupport.AssertRow(t, "add", data.Add(nil), InitData())

}

