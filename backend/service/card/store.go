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

func (s *Store) GetCardBySlug(slug string) (*types.CardResponse, error) {
	row := s.db.QueryRow("SELECT * FROM cards WHERE slug = $1", slug)

	c, err := scanRowIntoCard(row)

	if err != nil {
		return nil, err
	}

	rows, err := s.db.Query("SELECT id, card_id, card_item_id, name, value, created_at, updated_at FROM fields WHERE card_id = $1", c.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var fields []types.FieldResponse
	for rows.Next() {
		var field types.FieldResponse
		if err := rows.Scan(
			&field.ID,
			&field.CardID,
			&field.CardItemID,
			&field.Name,
			&field.Value,
			&field.CreatedAt,
			&field.UpdatedAt,
		); err != nil {
			return nil, err
		}
		fields = append(fields, field)
	}

	// Build the CardResponse
	cardResponse := &types.CardResponse{
		ID:        c.ID,
		Name:      c.Name,
		Slug:      c.Slug,
		UserID:    c.UserID,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
		CardItems: fields,
	}

	return cardResponse, nil
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

func (s *Store) CreateCardItemField(ctx context.Context, f types.CardItemField) error {
	query := "INSERT INTO fields(card_id, card_item_id, name, value) VALUES ($1, $2, $3, $4)"

	_, err := s.db.ExecContext(ctx, query, f.CardID, f.CardItemID, f.Name, f.Value)
	if err != nil {
		return err
	}

	return nil
}
