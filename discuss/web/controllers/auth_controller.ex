defmodule Discuss.AuthController do
  use Discuss.Web, :controller
  plug Ueberauth

  def callback(%{assigns: %{ueberauth_auth: auth}} = conn, params) do
    IO.puts "+++++++"
    IO.inspect(conn.assigns)
    IO.puts "+++++++"
    IO.inspect(params)
    IO.puts "+++++++"
    render conn, "index.html"
  end
end
