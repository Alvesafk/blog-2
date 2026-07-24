import { useState } from 'react'

export default function TextAreaForm({ url, onPost, inputPlaceholder, submitPlaceholder }) {
	const [formData, setFormData] = useState({
		author: "",
		content: "",
	})
	const [loading, setLoading] = useState(false)
	const [error, setError] = useState(null)

	function handleChange(e) {
		const { name, value } = e.target
		setFormData((prev) => ({ ...prev, [name]: value }))
	}

	async function handleSubmit(e) {
		e.preventDefault()
		setLoading(true)
		setError(null)

		try {
			const response = await fetch(url, {
				method: "POST",
				headers: {
					"Content-Type": "application/json"
				},
				body: JSON.stringify(formData)
			})

			if (!response.ok) {
				throw new Error(`Error on post: ${response.status}`)
			}

			onPost()
		} catch (err) {
			setError(err.message)
		} finally {
			setFormData({
				author: "",
				content: "",
			})
			setLoading(false)
		}
	}

	return (
		<form onSubmit={handleSubmit} className='container is-max-desktop my-2'>
			<div className='field'>
				<label className='label'>Author</label>
				<div className='control'>
					<input
						className='input'
						type="text"
						name="author"
						placeholder='Who are you?'
						value={formData.author}
						onChange={handleChange}
					/>
				</div>
			</div>
			<div className='field'>
				<label className='label'>Comment</label>
				<div className='control'>
					<textarea
						className='textarea'
						rows="3"
						cols="100"
						placeholder={inputPlaceholder}
						name='content'
						value={formData.content}
						onChange={handleChange}
					></textarea>
				</div>
			</div>
			<div className='field'>
				<div className='control'>
					<button className={`button is-link ${loading ? 'is-loading' : ''}`} type="submit" disabled={loading}>
						{loading ? "Posting..." : submitPlaceholder}
					</button>
				</div>
			</div>
			{error && <p>{error}</p>}
		</form>
	)
}
