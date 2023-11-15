import Image from 'next/image'

export default function Home() {
  return (
    <main className="flex min-h-screen flex-col items-center justify-between p-24">
      <div className="">

      <h1>Go-short</h1>
      <p> A URL shortener built in Golang</p>
      </div>

      <div>
   <form>
   <label>URL</label>
          <input className="border-2 px-2 py-2 rounded-md" type="text" name="name" />
          <button type="submit" className="border-2 px-2 py-2 rounded-md">Generate shorten URL</button>
    </form>
    </div>
     
    </main>
  )
}
