SET FOREIGN_KEY_CHECKS = 0;
DELETE FROM books WHERE title IN (
'To Kill a Mockingbird',
'1984',
'Pride and Prejudice',
'The Great Gatsby',
'Moby Dick',
'War and Peace',
'The Catcher in the Rye',
'The Hobbit',
'Brave New World',
'The Lord of the Rings'
);
SET FOREIGN_KEY_CHECKS=1;