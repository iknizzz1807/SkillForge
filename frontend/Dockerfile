FROM node:22-alpine

# Thiết lập thư mục làm việc
WORKDIR /app


# Copy package.json trước
COPY package.json ./

# Xóa package-lock.json nếu có (điều này giúp tránh xung đột)
RUN if [ -f package-lock.json ]; then rm package-lock.json; fi

# Cài đặt dependencies với các flag đặc biệt
RUN npm install --force

# Copy toàn bộ source code
COPY . .

# Thiết lập port
EXPOSE 5173

# Command để chạy ứng dụng trong development
CMD ["npm", "run", "dev", "--", "--host", "0.0.0.0"]