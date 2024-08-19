-- +goose Up
-- +goose StatementBegin
INSERT INTO card_items (name) VALUES
('Images'),
('Colors'),
('Name'),
('Job Title'),
('Department'),
('Company Name'),
('Email'),
('Phone'),
('Company URL'),
('Link'),
('Address'),
('X'),
('Instagram'),
('Threads'),
('LinkedIn'),
('Facebook'),
('YouTube'),
('Snapchat'),
('TikTok'),
('Twitch'),
('Yelp'),
('WhatsApp'),
('Signal'),
('Discord'),
('Skype'),
('Telegram'),
('Github'),
('Calendly'),
('PayPal'),
('Venmo'),
('CashApp');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM card_items;
-- +goose StatementEnd
