package kuji

import (
	"testing"
	"github.com/golang/mock/gomock"
	"reflect"
)

func TestNewKuji(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStrategy := NewMockKujiStrategy(ctrl)
	k := NewKuji(mockStrategy)

	instanceType := reflect.TypeOf(k)
	kujiType := reflect.TypeOf(Kuji{})
	if instanceType != kujiType {
		t.Error("TypeError, instance type mismatched!")
	}
	t.Log("Instance type matched with Kuji :)", instanceType, kujiType)
}

func TestKuji_PickOneByKey(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStrategy := NewMockKujiStrategy(ctrl)
	mockStrategy.EXPECT().PickOneByKey("PickOneByKey").Return("1", nil).Times(1)
	k := Kuji{
		strategy: mockStrategy,
	}

	str, err := k.PickOneByKey("PickOneByKey")
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
	k := Kuji{
		strategy: mockStrategy,
	}

	str, err := k.PickAndDeleteOneByKey("PickAndDeleteOneByKey")
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
	k := Kuji{
		strategy: mockStrategy,
	}
	str, err := k.PickOneByKeyAndIndex("PickOneByKeyAndIndex", int64(1))
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
			Id: 1,
			Weight: 100,
		},
	}
	mockStrategy := NewMockKujiStrategy(ctrl)
	mockStrategy.EXPECT().RegisterCandidatesWithKey("RegisterCandidatesWithKey", candidates).Return(int64(200), nil).Times(1)
	k := Kuji{
		strategy: mockStrategy,
	}
	success, err := k.RegisterCandidatesWithKey("RegisterCandidatesWithKey", candidates)
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
	k := Kuji{
		strategy: mockStrategy,
	}
	length, err := k.Len("Len")
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
	k := Kuji{
		strategy: mockStrategy,
	}
	length, err := k.List("List")
	if err != nil {
		t.Error("Something went wrong...")
	}
	if len(length) != len(elements) {
		t.Fatalf("Element length mismatched (expected: %d)", len(elements))
	}
	t.Logf("Result: %d", len(length))
}