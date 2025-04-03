FROM node:22-alpine

# Thiết lập thư mục làm việc
WORKDIR /app


# Copy package.json trước
COPY package.json ./

# Cài đặt dependencies với các flag đặc biệt
RUN npm install

# Copy toàn bộ source code
COPY . .

# Thiết lập port
EXPOSE 5173

# Command để chạy ứng dụng trong development
CMD ["npm", "run", "dev", "--", "--host", "0.0.0.0"]