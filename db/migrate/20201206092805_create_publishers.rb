class CreatePublishers < ActiveRecord::Migration[6.0]
  def change
    create_table :publishers do |t|
      t.string :name, limit: 30, null: false
      t.timestamps
    end

    add_reference :books, :publisher, foreign_key: true
  end
end
