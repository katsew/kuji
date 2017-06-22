package kuji

import (
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

func TestNewKuji(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	k := NewKuji()
	l := NewKuji()

	kType := reflect.TypeOf(k)
	lType := reflect.TypeOf(l)
	if kType != lType {
		t.Error("TypeError, instance type mismatched!")
	}
	t.Log("Instance type matched with Kuji :)", kType, lType)
}

func TestKuji_PickOneByKey(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStrategy := NewMockKujiStrategy(ctrl)
	mockStrategy.EXPECT().PickOneByKey("PickOneByKey").Return("1", nil).Times(1)
	k := NewKuji()
	k.Use("mock", KujiStrategyConfig{
		mockStrategy,
		nil,
	})

	str, err := k.PickOneByKey("mock", "PickOneByKey")
	if err != nil {
		t.Error("Something went wrong...")
	}
	t.Logf("Result: %s", str)
}

func TestKuji_PickAndDeleteOneByKey(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStrategy := NewMockKujiStrategy(ctrl)
	mockStrategy.EXPECT().PickAndDeleteOneByKey("PickAndDeleteOneByKey").Return("1", nil).Times(1)
	k := NewKuji()
	k.Use("mock", KujiStrategyConfig{
		mockStrategy,
		nil,
	})

	str, err := k.PickAndDeleteOneByKey("mock", "PickAndDeleteOneByKey")
	if err != nil {
		t.Error("Something went wrong...")
	}
	t.Logf("Result: %s", str)
}

func TestKuji_PickOneByKeyAndIndex(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStrategy := NewMockKujiStrategy(ctrl)
	mockStrategy.EXPECT().PickOneByKeyAndIndex("PickOneByKeyAndIndex", int64(1)).Return("1", nil).Times(1)
	k := NewKuji()
	k.Use("mock", KujiStrategyConfig{
		mockStrategy,
		nil,
	})
	str, err := k.PickOneByKeyAndIndex("mock", "PickOneByKeyAndIndex", int64(1))
	if err != nil {
		t.Error("Something went wrong...")
	}
	t.Logf("Result: %s", str)
}

func TestKuji_RegisterCandidatesWithKey(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	candidates := []KujiCandidate{
		KujiCandidate{
			Id:     1,
			Weight: 100,
		},
	}
	mockStrategy := NewMockKujiStrategy(ctrl)
	mockStrategy.EXPECT().RegisterCandidatesWithKey("RegisterCandidatesWithKey", candidates).Return(int64(200), nil).Times(1)
	k := NewKuji()
	k.Use("mock", KujiStrategyConfig{
		mockStrategy,
		nil,
	})
	success, err := k.RegisterCandidatesWithKey("mock", "RegisterCandidatesWithKey", candidates)
	if err != nil {
		t.Error("Something went wrong...")
	}
	t.Logf("Result: %d", success)
}

func TestKuji_Len(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStrategy := NewMockKujiStrategy(ctrl)
	mockStrategy.EXPECT().Len("Len").Return(int64(1000), nil).Times(1)
	k := NewKuji()
	k.Use("mock", KujiStrategyConfig{
		mockStrategy,
		nil,
	})
	length, err := k.Len("mock", "Len")
	if err != nil {
		t.Error("Something went wrong...")
	}
	t.Logf("Result: %d", length)
}

func TestKuji_List(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	elements := []string{"1", "2", "3", "4", "5"}
	mockStrategy := NewMockKujiStrategy(ctrl)
	mockStrategy.EXPECT().List("List").Return(elements, nil).Times(1)
	k := NewKuji()
	k.Use("mock", KujiStrategyConfig{
		mockStrategy,
		nil,
	})
	length, err := k.List("mock", "List")
	if err != nil {
		t.Error("Something went wrong...")
	}
	if len(length) != len(elements) {
		t.Fatalf("Element length mismatched (expected: %d)", len(elements))
	}
	t.Logf("Result: %d", len(length))
}
