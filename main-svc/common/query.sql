INSERT INTO users (id, name, username, password) VALUES (1, 'John Doe', 'johndoe', 'test123');

-- This query the link table return all quotes based on users.id 
-- SELECT q.id, q.body, q.author, q.quote_source FROM quotes q
-- JOIN userfavoritesquotes on (q.id = userfavoritesquotes.quotes_id)
-- JOIN users ON (users.id = userfavoritesquotes.user_id)
-- WHERE users.id = 1;