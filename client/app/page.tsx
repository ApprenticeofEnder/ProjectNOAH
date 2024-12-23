import { useQuery, gql } from '@apollo/client';

export default function Home() {
  const GET_CAMPAIGNS_FULL = gql`
    query {
      campaigns {
        title
        description
        user {
          email
          supabaseId
        }
        npcs {
          name
          notes
          combatClass {
            name
            role
          }
        }
        combats {
          title
          notes
        }
        scenes {
          title
          notes
        }
      }
    }
  `

  const { loading, error, data } = useQuery(GET_CAMPAIGNS_FULL);

  if (loading) return <p>Loading...</p>;

  if (error) return <p>Error : {error.message}</p>;

  console.log(data);

  return (
    <div>
      <h1>Under Construction</h1>
      <p>{`Project NOAH's front end is under construction! Check back soon.`}</p>
    </div>
  );
}
