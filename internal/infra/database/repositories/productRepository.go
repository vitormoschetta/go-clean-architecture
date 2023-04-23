package repositories

import (
	"database/sql"
	"log"

	"github.com/vitormoschetta/go/internal/domain/interfaces"
	"github.com/vitormoschetta/go/internal/domain/models"
)

type ProductRepository struct {
	Db *sql.DB
}

func NewProductRepository(db *sql.DB) interfaces.IProductRepository {
	return &ProductRepository{Db: db}
}

func (r *ProductRepository) FindAll() (products []models.Product, err error) {
	query := "SELECT p.id, p.name, p.price, c.id, c.name "
	query += "FROM products p "
	query += "INNER JOIN categories c ON p.category_id = c.id"
	rows, err := r.Db.Query(query)
	if err != nil {
		log.Print(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var p models.Product
		err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Category.ID, &p.Category.Name)
		if err != nil {
			log.Print(err)
			continue
		}
		products = append(products, p)
	}
	return
}

func (r *ProductRepository) FindByID(id string) (product models.Product, err error) {
	row := r.Db.QueryRow("SELECT id, name, price, category_id FROM products WHERE id = ?", id)
	err = row.Scan(&product.ID, &product.Name, &product.Price)
	if err != nil {
		log.Print(err)
	}
	return
}

func (r *ProductRepository) FindByCategory(categoryID string) (products []models.Product, err error) {
	rows, err := r.Db.Query("SELECT id, name, price, category_id FROM products WHERE category_id = ?", categoryID)
	if err != nil {
		log.Print(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var p models.Product
		err := rows.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil {
			log.Print(err)
			continue
		}
		products = append(products, p)
	}
	return
}

func (r *ProductRepository) Save(p models.Product) error {
	stmt, err := r.Db.Prepare("INSERT INTO products (id, name, price, category_id) VALUES (?, ?, ?, ?)")
	if err != nil {
		log.Print(err)
		return err
	}
	res, err := stmt.Exec(p.ID, p.Name, p.Price, p.Category.ID)
	if err != nil {
		log.Print(err)
		return err
	}
	if res != nil {
		log.Print("Product created")
	}
	return nil
}

func (r *ProductRepository) Update(p models.Product) error {
	stmt, err := r.Db.Prepare("UPDATE products SET name = ?, price = ?, category_id = ? WHERE id = ?")
	if err != nil {
		log.Print(err)
		return err
	}
	res, err := stmt.Exec(p.Name, p.Price, p.Category.ID, p.ID)
	if err != nil {
		log.Print(err)
		return err
	}
	if res != nil {
		log.Print("Product updated")
	}
	return nil
}

func (r *ProductRepository) Delete(id string) error {
	stmt, err := r.Db.Prepare("DELETE FROM products WHERE id = ?")
	if err != nil {
		log.Print(err)
		return err
	}
	res, err := stmt.Exec(id)
	if err != nil {
		log.Print(err)
		return err
	}
	if res != nil {
		log.Print("Product deleted")
	}
	return nil
}

func (r *ProductRepository) ApplyPromotionOnProductsByCategory(categoryId string, percentage float64) error {
	stmt, err := r.Db.Prepare("UPDATE products SET price = price - (price * ?) WHERE category_id = ?")
	if err != nil {
		log.Print(err)
		return err
	}
	res, err := stmt.Exec(percentage, categoryId)
	if err != nil {
		log.Print(err)
		return err
	}
	if res != nil {
		log.Print("Products updated")
	}
	return nil
}
