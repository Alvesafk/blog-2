import About from '../modules/About.jsx'
import Currently from '../modules/Currently.jsx'
import Gif from '../modules/Gif.jsx'
import Greet from '../modules/Greet.jsx'
import LatestPost from '../modules/LatestPost.jsx'
import Social from '../modules/Social.jsx'

export default function Home() {
	return (
		<>
			<Greet />
			<div className='container'>
				<div className='columns'>
					<Currently />
					<Social />
				</div>
			</div>
			<div>
				<LatestPost />
				<Gif url={"https://alvesafk.com/static/media/tux-linux.gif"} />
			</div>
			<About />
		</>
	);
}
