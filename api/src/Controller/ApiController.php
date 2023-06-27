<?php

namespace App\Controller;

use App\Repository\CountryRepository;
use Symfony\Bundle\FrameworkBundle\Controller\AbstractController;
// use Symfony\Component\HttpFoundation\Response;

// use Symfony\Bundle\FrameworkBundle\Controller\AbstractController;
use Symfony\Component\HttpFoundation\JsonResponse;
use Symfony\Component\Routing\Annotation\Route;

class ApiController extends AbstractController
{
    #[Route('/api/v1/countries', name: 'app_api')]
 public function index(CountryRepository $countryRepository): JsonResponse
    {
        $countries = $countryRepository->findAll();

        $data = [];
        foreach ($countries as $country) {
            $data[] = [
                'name' => $country->getNameOfficialFr(),
                'flag' => $country->getFlagUrl(),
                'population' => $country->getPopulation(),
            ];
        }

        return $this->json($data);
    }
}
