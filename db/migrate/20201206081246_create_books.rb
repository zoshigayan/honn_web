class CreateBooks < ActiveRecord::Migration[6.0]
  def change
    create_table :books do |t|
      t.string :title, limit: 100, null: false
      t.string :subtitle, limit: 100, null: false
      t.date :published_at, null: false
      t.timestamps
    end
  end
end
