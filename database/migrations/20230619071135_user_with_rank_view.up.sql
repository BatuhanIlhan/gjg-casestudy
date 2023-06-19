CREATE VIEW user_with_rank AS
SELECT *, RANK() OVER(ORDER BY points DESC) AS rank
FROM "user";