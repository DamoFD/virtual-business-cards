package card

import (
	"context"
	"database/sql"

	"github.com/DamoFD/virtual-business/types"
)

// Store is a struct that contains a database connection.
// It contains methods for getting and creating users.
type Store struct {
	db *sql.DB
}

// NewStore creates a new user store.
// It takes a database connection as a parameter.
// It returns a pointer to the Store struct.
func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func scanRowsIntoCard(row *sql.Rows) (*types.Card, error) {
	c := new(types.Card)
	err := row.Scan(
		&c.ID,
		&c.UserID,
		&c.Name,
		&c.Slug,
		&c.CreatedAt,
		&c.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func scanRowIntoCard(row *sql.Row) (*types.Card, error) {
	c := new(types.Card)
	err := row.Scan(
		&c.ID,
		&c.UserID,
		&c.Name,
		&c.Slug,
		&c.CreatedAt,
		&c.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (s *Store) GetCardBySlug(slug string) (*types.Card, error) {
	row := s.db.QueryRow("SELECT * FROM cards WHERE slug = $1", slug)

	c, err := scanRowIntoCard(row)

	if err != nil {
		return nil, err
	}

	return c, nil
}

func (s *Store) GetCardsByUserID(userID int) ([]*types.Card, error) {

	rows, err := s.db.Query("SELECT * FROM cards WHERE user_id = $1", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cards := []*types.Card{}
	for rows.Next() {
		c, err := scanRowsIntoCard(rows)
		if err != nil {
			return nil, err
		}
		cards = append(cards, c)
	}

	return cards, nil
}

func (s *Store) CreateCard(ctx context.Context, c types.Card) (types.Card, error) {
	var card types.Card

	query := "INSERT INTO cards(user_id, name, slug) VALUES($1, $2, $3) RETURNING id, user_id, name, slug, created_at, updated_at"

	err := s.db.QueryRow(query, c.UserID, c.Name, c.Slug).Scan(
		&card.ID,
		&card.UserID,
		&card.Name,
		&card.Slug,
		&card.CreatedAt,
		&card.UpdatedAt,
	)

	if err != nil {
		return card, err
	}

	return card, nil
}
func (s *Store) UpdateCard(ctx context.Context, card types.Card) (types.Card, error) {
	return types.Card{}, nil
}
func (s *Store) DeleteCard(ctx context.Context, cardID int) error {
	return nil
}
