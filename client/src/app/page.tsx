import HotelCardList from "@/components/home/hotel-list";
import Destination from "@/components/home/destination";
import Hero from "@/components/home/hero";
import Promotion from "@/components/home/promotion";
import About from "@/components/home/about";
import Testmonial from "@/components/home/testmonial";


export default async function Home() {

  return (
    <section className="max-w-7xl mx-auto px-4">
      <Hero />
      <About />
      <Destination />
      <HotelCardList />
      <Promotion />
      <Testmonial />
    </section>
  );
}
