<?php

namespace App\Entity;

use App\Repository\CountryRepository;
use Doctrine\ORM\Mapping as ORM;

#[ORM\Entity(repositoryClass: CountryRepository::class)]
#[ORM\Table(name:"countries")]
class Country
{
    #[ORM\Id]
    #[ORM\GeneratedValue]
    #[ORM\Column]
    private ?int $id = null;

    #[ORM\Column(length: 255)]
    private ?string $name_official_fr = null;

    #[ORM\Column(length: 255)]
    private ?string $flag_url = null;

    #[ORM\Column]
    private ?int $population = null;

    public function getId(): ?int
    {
        return $this->id;
    }

    public function getNameOfficialFr(): ?string
    {
        return $this->name_official_fr;
    }

    public function setNameOfficialFr(string $name_official_fr): static
    {
        $this->name_official_fr = $name_official_fr;

        return $this;
    }

    public function getFlagUrl(): ?string
    {
        return $this->flag_url;
    }

    public function setFlagUrl(string $flag_url): static
    {
        $this->flag_url = $flag_url;

        return $this;
    }

    public function getPopulation(): ?int
    {
        return $this->population;
    }

    public function setPopulation(int $population): static
    {
        $this->population = $population;

        return $this;
    }
}
