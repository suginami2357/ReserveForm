package places

type Repository interface {
	Index() []Place
}
