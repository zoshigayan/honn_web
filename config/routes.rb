Rails.application.routes.draw do
  get "/", to: "books#index"
  resources :books, only: [:show]

  resources :authors, only: [:index, :show]
  resources :publishers, only: [:index, :show]
end
