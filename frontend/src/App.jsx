import { BrowserRouter, Routes, Route } from "react-router";
import MainLayout from "./layouts/MainLayout";
import Home from "./pages/Home";
import Resources from "./pages/Resources";
import Guestbook from "./pages/Guestbook";
import Posts from "./pages/Posts";
import Post from "./pages/Post";
import NotFound from "./pages/404";

function App() {
	return (
		<>
			<BrowserRouter>
				<Routes>
					<Route element={<MainLayout />}>
						<Route path="/" element={<Home />} />
						<Route path="/resources" element={<Resources />} />
						<Route path="/guestbook" element={<Guestbook />} />
						<Route path="/posts" element={<Posts />} />
						<Route path="/posts/:name" element={<Post />} />
						<Route path="*" element={<NotFound />} />
					</Route>
				</Routes>
			</BrowserRouter>
		</>
	);
}
export default App;
