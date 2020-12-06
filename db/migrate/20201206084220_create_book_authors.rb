class CreateBookAuthors < ActiveRecord::Migration[6.0]
  def change
    create_table :book_authors do |t|
      t.references :book, index: true, null: false, foreign_key: true
      t.references :author, index: true, null: false, foreign_key: true

      t.timestamps
    end
  end
end
