import Head from 'next/head'
import Link from 'next/link'
import Image from 'next/image'

export default function Home() {
  return (
    <div>
      <Head>
        <title>Edison Aguiar</title>
        <meta name="keywords" content="Edison Aguiar,Edison Aguiar de Souza Neto" />
        <meta name="author" content="Edison Aguiar" />
      </Head>

      <main>
        <Image
          src="/images/profile.jpg"
          height={130}
          width={130}
          alt="Edison Aguiar"
        />

        <h1 className="title">
          Edison Aguiar
        </h1>

        <div>
          <p>
            I'm a software engineer, my primary focus now is learning <a href="https://www.rust-lang.org/">rust</a> and build some systems with it.
          </p>
          A list of things I'm building or will build:
          <ul>
            <li>HTTP Server</li>
            <li>DNS Server</li>
            <li>File System</li>
            <li>Distributed Commmit Log</li>
            <li>Database based on <a href="https://hstore.cs.brown.edu/">H-Store</a></li>
            <li>Distributed In-Memory Cache</li>
          </ul>

          <h2>Main interests</h2>
          <p>Databases, Distributed Systems, Astronomy, Eletronics.</p>

          <h2>Contact/Info</h2>
          <ul>
            <li>
              <Link href="https://github.com/eaneto">
                <a>Github</a>
              </Link>
            </li>
            <li>
              <Link href="mailto:edison.aguiar.neto@gmail.com">
                <a>edison.aguiar.neto@gmail.com</a>
              </Link>
            </li>
          </ul>

          <h2>Research</h2>
          <p>
            At my university I've done research in solar physics and databases, specifically PostgreSQL.
            I published some papers from that research:
          </p>
          <ul>
            <li>
              <p>
                Tracking Automation Of The 7 GHZ Solar Radio Polarimeter Using a Paramount MEII Robotic Equatorial Mount
              </p>
            </li>
            <li>
              <p>
                Performance of serializable transactions in high concurrency scenarios in PostgreSQL (Portuguese)
              </p>
            </li>
          </ul>
        </div>
      </main>
    </div>
  )
}
