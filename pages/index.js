import Head from 'next/head'
import Link from 'next/link'
import Image from 'next/image'
import homeStyle from '../styles/home.module.scss'

export default function Home() {
  return (
    <div className={homeStyle.container}>
      <Head>
        <title>Edison Aguiar</title>
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
            Software Engineer
          </p>
          <table>
            <tr>
              <th>Github</th>
              <td>
                <Link href="https://github.com/eaneto">
                  <a>https://github.com/eaneto</a>
                </Link>
              </td>
            </tr>
            <tr>
              <th>Email</th>
              <td>
                <Link href="mailto:edison.aguiar.neto@gmail.com">
                  <a>edison.aguiar.neto@gmail.com</a>
                </Link>
              </td>
            </tr>
          </table>
        </div>
      </main>
    </div>
  )
}
