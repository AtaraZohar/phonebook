# שימוש ב-Base Image של Go, גרסה קלה (Alpine)
FROM golang:1.20-alpine

# הגדרת תיקיית העבודה בפרויקט
WORKDIR /app

# יצירת תיקיית הלוגים
RUN mkdir -p /app/logs

# העתקת קבצי ההגדרות עבור Go Modules
COPY go.mod ./
COPY go.sum ./

# הורדת כל התלויות
RUN go mod download
RUN go mod tidy

# העתקת כל הקבצים מתיקיית הפרויקט
COPY . .

# בניית הקובץ הבינארי בשם phonebook-api
RUN go build -o /phonebook-api

# חשיפת פורט 8080 לצורך גישה ל-API
EXPOSE 8080

# פקודת הרצה של הקובץ הבינארי phonebook-api
CMD ["/phonebook-api"]
