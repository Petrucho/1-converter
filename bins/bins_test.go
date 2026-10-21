package bins

import (
	"testing"
	"time"
)

// fakeRepo — мок хранилища в памяти. Именно ради таких подстановок
// Service зависит от интерфейса Repository, а не от storage.Storage.
type fakeRepo struct {
	saved *BinList
}

func (f *fakeRepo) SaveBins(list *BinList) error {
	f.saved = list
	return nil
}

func (f *fakeRepo) ReadBins() (*BinList, error) {
	if f.saved == nil {
		return NewBinList(), nil
	}
	return f.saved, nil
}

func TestService_CreateAndSave(t *testing.T) {
	repo := &fakeRepo{}

	service, err := NewService(repo)
	if err != nil {
		t.Fatalf("NewService: %v", err)
	}

	if _, err := service.Create("bin-001", true, time.Now(), "Private Bin"); err != nil {
		t.Fatalf("Create: %v", err)
	}
	if err := service.Save(); err != nil {
		t.Fatalf("Save: %v", err)
	}

	if repo.saved == nil || len(repo.saved.Bins) != 1 {
		t.Fatalf("ожидалась одна сохранённая корзина, получили %+v", repo.saved)
	}
	if repo.saved.Bins[0].Id != "bin-001" {
		t.Errorf("Id = %q, ожидался bin-001", repo.saved.Bins[0].Id)
	}
}

func TestService_Load(t *testing.T) {
	repo := &fakeRepo{saved: NewBinList()}
	repo.saved.Bins = append(repo.saved.Bins, Bin{Id: "bin-002", Name: "Public", CreatedAt: time.Now()})

	service, err := NewService(repo)
	if err != nil {
		t.Fatalf("NewService: %v", err)
	}

	list, err := service.Load()
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if len(list.Bins) != 1 || list.Bins[0].Id != "bin-002" {
		t.Fatalf("неожиданный список: %+v", list)
	}
}

func TestNewService_NilRepo(t *testing.T) {
	if _, err := NewService(nil); err == nil {
		t.Fatal("ожидалась ошибка при nil-репозитории")
	}
}

func TestCreateBin_Validation(t *testing.T) {
	now := time.Now()
	cases := []struct {
		name string
		id   string
		at   time.Time
		bin  string
		want error
	}{
		{"пустой id", "", now, "name", ErrInvalidID},
		{"нулевое время", "id", time.Time{}, "name", ErrInvalidCreatedAt},
		{"пустое имя", "id", now, "", ErrInvalidName},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if _, err := CreateBin(c.id, false, c.at, c.bin); err != c.want {
				t.Errorf("err = %v, ожидалась %v", err, c.want)
			}
		})
	}
}
