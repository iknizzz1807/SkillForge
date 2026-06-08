#!/bin/bash
sed -i '/func (r \*ChatRepository) DeleteGroupMessages/,$d' backend/internal/repositories/chat_repository.go
sed -i '/func (r \*TaskRepository) DeleteTasksByProjectID/,$d' backend/internal/repositories/task_repository.go
