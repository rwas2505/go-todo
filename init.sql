CREATE TABLE IF NOT EXISTS taskCategories(
taskCategoryId SERIAL PRIMARY KEY,
taskCategoryName VARCHAR(100) NOT NULL
);

CREATE TABLE IF NOT EXISTS tasks(
taskId SERIAL PRIMARY KEY,
taskName VARCHAR(100) NOT NULL,
taskDescription TEXT,
taskCreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
taskIsComplete BOOLEAN DEFAULT FALSE,
taskCategoryId INT,
CONSTRAINT fk_taskCategory
	FOREIGN KEY(taskCategoryId)
	REFERENCES taskCategories(taskCategoryId)
);

INSERT INTO taskCategories(taskCategoryName)
SELECT 'House Chores'
WHERE
NOT EXISTS (
SELECT taskCategoryName FROM taskCategories WHERE taskCategoryName = 'House Chores'
);

INSERT INTO tasks
(taskName, taskDescription, taskIsComplete, taskCategoryId)
SELECT
'Mow Lawn', 'Do the front and back yards', false, (SELECT taskCategoryId FROM taskCategories WHERE taskCategoryName = 'House Chores')
WHERE NOT EXISTS (
SELECT taskName FROM tasks
WHERE taskName = 'Mow Lawn' AND taskDescription = 'Do the front and back yards' AND taskCategoryId = (SELECT taskCategoryId FROM taskCategories WHERE taskCategoryName = 'House Chores')
)