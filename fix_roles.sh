sed -i -e '/if role != "business" {/,/return/d' backend/internal/handlers/applications.go
sed -i -e '/if c.GetString("role") != "business" {/,/return/d' backend/internal/handlers/talentpool.go
sed -i -e '/if role != "business" {/,/return/d' backend/internal/handlers/business_info.go
